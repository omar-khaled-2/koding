module.exports = (options = {}, callback)->
  encoder = require 'htmlencode'

  options.intro                 ?= no
  options.landing               ?= no
  options.client               or= {}
  options.client.context       or= {}
  options.client.context.group or= "koding"
  options.client.connection    or= {}

  {argv} = require 'optimist'

  prefetchedFeeds = {}
  campaignData    = {}
  currentGroup    = {}
  {bongoModels, client, intro, landing, slug} = options

  createHTML = ->
    replacer            = (k, v)-> if 'string' is typeof v then encoder.XSSEncode v else v
    encodedFeed         = JSON.stringify prefetchedFeeds, replacer
    encodedCampaignData = JSON.stringify campaignData, replacer
    currentGroup        = JSON.stringify currentGroup, replacer
    landingOptions      = page : landing

    if client.connection?.delegate?.profile?.nickname
      {connection: {delegate}} = client
      {profile   : {nickname}} = delegate
      landingOptions.username  = nickname if delegate.type is "registered"

    landingOptions = JSON.stringify landingOptions
    """
    <script>
      console.time("Framework loaded");
      console.time("Koding.com loaded");
    </script>

    <!-- MIXPANEL -->
    <script>(function(e,b){if(!b.__SV){var a,f,i,g;window.mixpanel=b;a=e.createElement("script");a.type="text/javascript";a.async=!0;a.src=("https:"===e.location.protocol?"https:":"http:")+'//cdn.mxpnl.com/libs/mixpanel-2.2.min.js';f=e.getElementsByTagName("script")[0];f.parentNode.insertBefore(a,f);b._i=[];b.init=function(a,e,d){function f(b,h){var a=h.split(".");2==a.length&&(b=b[a[0]],h=a[1]);b[h]=function(){b.push([h].concat(Array.prototype.slice.call(arguments,0)))}}var c=b;"undefined"!==typeof d?c=b[d]=[]:d="mixpanel";c.people=c.people||[];c.toString=function(b){var a="mixpanel";"mixpanel"!==d&&(a+="."+d);b||(a+=" (stub)");return a};c.people.toString=function(){return c.toString(1)+".people (stub)"};i="disable track track_pageview track_links track_forms register register_once alias unregister identify name_tag set_config people.set people.set_once people.increment people.append people.track_charge people.clear_charges people.delete_user".split(" ");for(g=0;g<i.length;g++)f(c,i[g]);b._i.push([a,e,d])};b.__SV=1.2}})(document,window.mixpanel||[]);mixpanel.init("#{KONFIG.mixpanel}");</script>

    <script>KD.campaign={}</script>
    <script src='/a/js/kd.#{KONFIG.version}.js'></script>
    #{if intro then "<script src='/a/js/introapp.#{ KONFIG.version }.js'></script>" else ''}
    <script src='/a/js/koding.#{KONFIG.version}.js'></script>
    #{if landing then "<script src='/a/js/landingapp.#{ KONFIG.version }.js'></script>" else ''}
    <script>KD.prefetchedFeeds=#{encodedFeed};</script>
    <script>KD.currentGroup=#{currentGroup};</script>


    <!-- GOOGLE ANALYTICS -->
    <script>
      var _gaq = _gaq || [];
      _gaq.push(['_setAccount', 'UA-6520910-8']);
      _gaq.push(['_setDomainName', 'koding.com']);
      _gaq.push(['_trackPageview']);
      (function() {
        var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
        ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
        var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
      })();
    </script>

    <!-- ROLLBAR -->
    <script>
      var startTime = new Date().getTime();
      var _rollbarParams = {
        "server.environment": "production",
        "client.javascript.source_map_enabled": true,
        "client.javascript.code_version": "#{KONFIG.version}",
        "client.javascript.guess_uncaught_frames": true,
        checkIgnore: function(msg, file, line, col, err) {
          if ((new Date().getTime() - startTime) > 1000*60*60) {
            // ignore errors after the page has been open for 1hr
            return true;
          }
          return false;
        }
      };
      _rollbarParams["notifier.snippet_version"] = "2"; var _rollbar=["#{KONFIG.rollbar}", _rollbarParams]; var _ratchet=_rollbar;
      (function(w,d){w.onerror=function(e,u,l){_rollbar.push({_t:'uncaught',e:e,u:u,l:l});};var i=function(){var s=d.createElement("script");var
      f=d.getElementsByTagName("script")[0];s.src="//d37gvrvc0wt4s1.cloudfront.net/js/1/rollbar.min.js";s.async=!0;
      f.parentNode.insertBefore(s,f);};if(w.addEventListener){w.addEventListener("load",i,!1);}else{w.attachEvent("onload",i);}})(window,document);
    </script>
    #{if argv.t then "<script src=\"/a/js/tests.js\"></script>" else ''}
    """

  generateScript = ->
    bongoModels.JReferral.isCampaingValid (err, status, campaign)->
      if not err
        content = campaign?.content
        campaignData = {status, content}



      kallback = ->
        html = createHTML()
        callback null, html

      # get group logo and coverphoto
      if slug # if not, it is koding
        bongoModels.JGroup.one {slug}, (err, group) ->
          console.log err if err
          if group
            currentGroup =
              logo       : group.customize?.logo or ""
              coverPhoto : group.customize?.coverPhoto or ""
              id         : group.getId()
        kallback()
      else
        kallback()



  {delegate} = options.client.connection
  # if user is exempt or super-admin do not cache his/her result set
  return generateScript()  if delegate and delegate.checkFlag ['super-admin', 'exempt']

  Cache  = require '../cache/main'
  feedFn = require '../cache/feed'

  getCacheKey =-> return "scriptblock-#{options.client.context.group}"

  Cache.fetch getCacheKey(), feedFn, options, (err, data)->
    prefetchedFeeds = data    # this is updating the prefetchedFeeds property
    return generateScript()   # we can generate html here
