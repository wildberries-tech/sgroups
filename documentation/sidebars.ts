import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';

/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */
const sidebars: SidebarsConfig = {

  informationSidebar: [
    {
      type: 'doc',
      label: 'Введение',
      id: 'info/introduction',
    },
    {
      type: 'doc',
      label: 'Выбор инструмента',
      id: 'info/toolset'
    },
    {
      type: 'category',
      label: 'Архитектура',
      items: [{
        type: 'doc',
        label: 'Абстракции',
        id: 'info/architecture/abstractions'
      },
      {
        type: 'doc',
        label: 'Компоненты',
        id: 'info/architecture/components'
      }]
    }
  ],

  settingsSidebar: [{
    type: 'category',
    label: 'Установка',
    items: [{
      type: 'doc',
      label: 'Системные требования',
      id: 'settings/installation/installation-system-requirements'
    },
    {
      type: 'doc',
      label: 'Сервер',
      id: 'settings/installation/installation-server'
    },
    {
      type: 'doc',
      label: 'Агент',
      id: 'settings/installation/installation-agent'
    },
    {
      type: 'doc',
      label: 'Terraform',
      id: 'settings/installation/installation-terraform'
    }]
  },
  {
    type: 'doc',
    label: 'Как пользоваться',
    id: 'settings/how-to-use'
  },
  {
    type: 'doc',
    label: 'Отладка',
    id: 'settings/debugging'
  }],

  docsForUsers: [{
    type: 'doc',
    label: 'Собственные Security Group',
    id: 'user/swarm-spec'
  },
  {
    type: 'category',
    label: 'Terraform',
    items: [{
      type: 'doc',
      label: 'resource-networks',
      id: 'user/terraform/resource-networks'
    },
    {
      type: 'doc',
      label: 'recource-rule-classic-s2f',
      id: 'user/terraform/recource-rule-classic-s2f'
    },
    {
      type: 'doc',
      label: 'recource-rule-classic-s2s-icmp',
      id: 'user/terraform/recource-rule-classic-s2s-icmp'
    },
    {
      type: 'doc',
      label: 'recource-rule-classic-s2s-tcp-udp',
      id: 'user/terraform/recource-rule-classic-s2s-tcp-udp'
    },
    {
      type: 'doc',
      label: 'recource-rule-ie-s2c-icmp',
      id: 'user/terraform/recource-rule-ie-s2c-icmp'
    },
    {
      type: 'doc',
      label: 'recource-rule-ie-s2c-tcp-udp',
      id: 'user/terraform/recource-rule-ie-s2c-tcp-udp'
    },
    {
      type: 'doc',
      label: 'recource-rule-ie-s2s-icmp',
      id: 'user/terraform/recource-rule-ie-s2s-icmp'
    },
    {
      type: 'doc',
      label: 'recource-rule-ie-s2s-tcp-udp',
      id: 'user/terraform/recource-rule-ie-s2s-tcp-udp'
    },
    {
      type: 'doc',
      label: 'recource-sgroup',
      id: 'user/terraform/recource-sgroup'
    }]
  }],

  toNft: [{
    type: 'doc',
    label: 'Мониторинг',
    id: 'to-nft/monitoring'
  },
  {
    type: 'doc',
    label: 'Конфигурация  правил',
    id: 'to-nft/rules-configuration'
  }],

  sgroups: [{
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
};

export default sidebars;
