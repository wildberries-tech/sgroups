export const sgroups: any = [{
    type: 'doc',
    label: 'Мониторинг',
    id: 'sgroups/monitoring'
  }, 
  {
    type: 'doc',
    label: 'Описание базы данных',
    id: 'sgroups/database'
  },
  {
    type: 'category',
    label: 'API',
    items: [{
      type: 'doc',
      label: 'POST /v1/sync',
      id: 'sgroups/api/v1/sync'
    },
    {
      type: 'doc',
      label: 'POST /v1/list/security-groups',
      id: 'sgroups/api/v1/security-groups'
    },
    {
      type: 'doc',
      label: 'GET /v1/{address}/sg',
      id: 'sgroups/api/v1/address-sg'
    },
    {
      type: 'doc',
      label: 'POST /v1/list/networks',
      id: 'sgroups/api/v1/networks'
    },
    {
      type: 'doc',
      label: 'GET /v1/sg/{sgName}/subnets',
      id: 'sgroups/api/v1/subnets'
    },
    {
      type: 'doc',
      label: 'POST /v1/sg-sg-icmp/rules',
      id: 'sgroups/api/v1/sg-sg-icmp-rules'
    },
    {
      type: 'doc',
      label: 'POST /v1/sg-icmp/rules',
      id: 'sgroups/api/v1/sg-icmp-rules'
    },
    {
      type: 'doc',
      label: 'POST /v1/rules',
      id: 'sgroups/api/v1/rules'
    },
    {
      type: 'doc',
      label: 'POST /v1/fqdn/rules',
      id: 'sgroups/api/v1/fqdn-rules'
    },
    {
      type: 'doc',
      label: 'POST /v1/cidr-sg/rules',
      id: 'sgroups/api/v1/cidr-sg-rules'
    },
    {
      type: 'doc',
      label: 'POST v1/ie-sg-sg/rules',
      id: 'sgroups/api/v1/ie-sg-sg-rules'
    },
    {
      type: 'doc',
      label: 'GET /v1/sync/status',
      id: 'sgroups/api/v1/status'
    }]
  }]