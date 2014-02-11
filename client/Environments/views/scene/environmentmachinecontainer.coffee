class EnvironmentMachineContainer extends EnvironmentContainer

  constructor:(options={}, data)->

    options.cssClass   = 'machines'
    options.itemClass  = EnvironmentMachineItem
    options.title      = 'Virtual Machines'
    options.itemHeight = 50

    super options, data

  loadItems:->
    super

    vmc = KD.getSingleton 'vmController'

    {entryPoint} = KD.config
    cmd = if entryPoint then 'fetchGroupVMs' else 'fetchVmNames'

    vmc[cmd] yes, (err, vms)=>
      if err or vms.length is 0
        @emit "DataLoaded"
        return warn "Failed to fetch VMs", err  if err
      addedCount = 0
      vms.forEach (vm)=>
        @addItem
          title     : vm
          cpuUsage  : KD.utils.getRandomNumber 100
          memUsage  : KD.utils.getRandomNumber 100
          activated : yes
        addedCount++
        @emit "DataLoaded"  if addedCount is vms.length