
export const toNft: any =  [{
    type: 'doc',
    label: 'Мониторинг',
    id: 'to-nft/monitoring'
  },
  {
    type: 'category',
    label: 'Конфигурация  правил',
    link: {
      type: 'doc',
      id: 'to-nft/rule-configuration/main'
    },
    items: [{
      type: 'doc',
      label: 'Config Base Rules',
      id: 'to-nft/rule-configuration/config-base-rules'
    },
    {
      type:'doc',
      label: 'Sgroup to Sgroup (tcp|udp)',
      id: 'to-nft/rule-configuration/s2s-tcp-udp'
    },
    {
      type: 'doc',
      label: 'Sgroup to Sgroup (icmp)',
      id: 'to-nft/rule-configuration/s2s-icmp'
    },
    {
      type: 'doc',
      label: 'Sgroup Default Rule (all)',
      id: 'to-nft/rule-configuration/default-rule-all'
    },
    {
      type: 'doc',
      label: 'Sgroup to Sgroup (ie: tcp|udp)',
      id: 'to-nft/rule-configuration/s2s-ie-tcp-udp'
    },
    {
      type: 'doc',
      label: 'Sgroup Default  Rule (icmp)',
      id: 'to-nft/rule-configuration/default-rule-icmp'
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