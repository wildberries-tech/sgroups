
export const toNft: any =  [{
    type: 'doc',
    label: 'Мониторинг',
    id: 'to-nft/monitoring'
  },
  {
    type: 'doc',
    label: 'Config Base Rules',
    id: 'to-nft/rule-configuration/config-base-rules'
  },
  {
    type: 'category',
    label: 'Конфигурация  ресурсов',
    link: {
      type: 'doc',
      id: 'to-nft/rule-configuration/main'
    },
    items: [{
      type: 'doc',
      label: 'Security Groups',
      id: 'to-nft/rule-configuration/default-rule-all'
    },
    {
      type:'doc',
      label: 'Sgroup to Sgroup',
      id: 'to-nft/rule-configuration/s2s'
    },
    {
      type: 'doc',
      label: 'Sgroup to Sgroup (ie: tcp|udp)',
      id: 'to-nft/rule-configuration/s2s-ie'
    }, 
    {
      type:'doc',
      label: 'Sgroup to CIDR (ie: tcp|udp)',
      id: 'to-nft/rule-configuration/s2c-ie-tcp-udp'
    },
    {
      type:'doc',
      label: 'Sgroup to FQDN (e: tcp|udp)',
      id: 'to-nft/rule-configuration/s2f-e-tcp-udp'
    }]
  }]