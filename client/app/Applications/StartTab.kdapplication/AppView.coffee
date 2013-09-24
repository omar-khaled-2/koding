class StartTabMainView extends JView

  constructor:(options = {}, data)->

    options.cssClass = 'start-tab'

    super options, data

    @appIcons   = {}
    appStorages = KD.getSingleton('appStorageController')
    @appStorage = appStorages.storage 'Finder', '1.1'
    @_connectAppsController()

    # Main view elements

    @loader = new KDLoaderView size : width : 16

    @appItemContainer = new StartTabAppItemContainer
      cssClass : 'app-item-container'
      delegate : this

    @serversWrapper = new KDView
      cssClass : 'file-container'

  # Application Specific Operations

  _connectAppsController:->
    @appsController = KD.getSingleton("kodingAppsController")

    @appsController.on "AppsRefreshed",   (apps)=> @decorateApps apps
    @appsController.on "AppsDataChanged", @bound "updateAppIcons"
    @appsController.on "InvalidateApp",   @bound "removeAppIcon"
    @appsController.on "UpdateAppData",   @bound "createAppIcon"

    @on 'refreshAppsMenuItemClicked', =>
      @appsController.syncAppStorageWithFS yes
    @on 'makeANewAppMenuItemClicked', =>
      @appsController.makeNewApp()

  addRealApps:->

    @removeAppIcons()
    @showLoader()
    @appsController.fetchApps (err, apps)=>
      if not @appsController._loadedOnce and apps and Object.keys(apps).length > 0
        @appsController.syncAppStorageWithFS()
      @decorateApps apps

  decorateApps:(apps)->

    apps or= @appsController.getManifests()

    @removeAppIcons()
    @showLoader()

    @appsController.appStorage.fetchValue 'shortcuts', (shortcuts)=>

      for own shortcut, manifest of shortcuts
        do (shortcut, manifest)=>
          @appItemContainer.addSubView @appIcons[manifest.name] = new AppShortcutButton
            delegate : @
          , manifest

      @createAllAppIcons apps
      @createGetMoreAppsButton()

      @hideLoader()

  createGetMoreAppsButton:->
    @appIcons['GET_MORE_APPS']?.destroy()
    @appItemContainer.addSubView @appIcons['GET_MORE_APPS'] = new GetMoreAppsButton
      delegate : @

  removeAppIcon:(appName)->
    appIcon = @appIcons[appName]
    return  unless appIcon
    appIcon.destroy()
    delete @appIcons[appName]

  removeAppIcons:->

    @appItemContainer.destroySubViews()
    @appIcons = {}

  updateAppIcons:(changes)->

    {removedApps, newApps, existingApps, force} = changes
    return @decorateApps()  if force or existingApps.length is 0
    @removeAppIcon app  for app in removedApps
    @createAllAppIcons @appsController.getManifests()  if newApps.length > 0

  createAppIcon:(app, appData, bulk=no)->

    appData or= @appsController.getManifest app
    return  unless appData

    oldIcon = @appIcons[app]
    @appItemContainer.addSubView newIcon = new StartTabAppThumbView
      delegate : this
    , appData

    if oldIcon
      newIcon.$().insertAfter oldIcon.$()
      oldIcon.destroy()

    @appIcons[app] = newIcon

    # To make sure its always the last icon
    @createGetMoreAppsButton()  unless bulk

  createAllAppIcons:(apps)->
    for own app, appData of apps
      do (app, appData)=>
        @createAppIcon app, appData, yes

    # To make sure its always the last icon
    @createGetMoreAppsButton()

  # Guest Notifications if necessary
  #
  # We need to inform Guest users to register a new account in 20 min.

  startGuestTimer:->
    return  unless KD.isGuest()
    unless $.cookie "guestForFirstTime"
      @utils.wait 5*60*1000, =>
        @showGuestNotification()
        $.cookie "guestForFirstTime", yes
    else
      @showGuestNotification()

  showGuestNotification: (guestTimeout = 20)->
    return  unless KD.isGuest()
    guestCreate      = new Date KD.whoami().meta.createdAt
    guestCreateTime  = guestCreate.getTime()

    endTime       = new Date(guestCreateTime + guestTimeout*60*1000)

    notification  = new GlobalNotification
      title       : "Your session will end in"
      targetDate  : endTime
      endTitle    : "Your session end, logging out."
      content     : "You can use Koding for 20 minutes without registering. <a href='/Register'>Register now</a>."
      callback    : =>
        return  unless KD.isGuest()
        {defaultVmName} = KD.getSingleton "vmController"
        KD.remote.api.JVM.removeByHostname defaultVmName, (err)->
          KD.getSingleton("finderController").unmountVm defaultVmName
          KD.getSingleton("vmController").emit 'VMListChanged'
          $.cookie "clientId", erase: yes
          $.cookie "guestForFirstTime", erase: yes

  # Common parts

  showLoader:->

    @loader.show()
    @$('h1.loaded, h2.loaded').addClass "hidden"
    @$('h2.loader').removeClass "hidden"

  hideLoader:->

    @loader.hide()
    @$('h2.loader').addClass "hidden"
    @$('h1.loaded, h2.loaded').removeClass "hidden"

  viewAppended:->

    super
    @addRealApps()
    @startGuestTimer()

  pistachio:->
    """
    <div class='app-list-wrapper'>
      <header>
        <h1 class="start-tab-header loaded hidden">This is your Development Area</h1>
        <h2 class="loaded hidden">You can install more apps on Apps section, or use the ones below that are already installed.</h2>
        <h2 class="loader">{{> @loader}} Loading applications...</h1>
      </header>
      {{> @appItemContainer}}
    </div>
    <div class='start-tab-split-options expanded'>
      <h3>Start with a workspace</h3>
    </div>
    <div class='start-tab-recent-container'>
      <h3>Your servers:</h3>
      {{> @serversWrapper}}
    </div>
    """
