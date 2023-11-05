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
  }]
};

export default sidebars;
