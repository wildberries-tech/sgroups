import React, { FC } from 'react'
import { VERSION } from '@site/src/constants/githubFacts'

export const GithubLink: FC = () => (
  <a
    href="https://github.com/H-BF/sgroups"
    target="_blank"
    rel="noopener noreferrer"
    className="navbar__item navbar__link header-github-link"
    aria-label="GitHub repository"
  >
    H-BF/sgroups
    <svg
      width="13.5"
      height="13.5"
      aria-hidden="true"
      viewBox="0 0 24 24"
      className="iconExternalLink_node_modules-@docusaurus-theme-classic-lib-theme-Icon-ExternalLink-styles-module"
    >
      <path
        fill="currentColor"
        d="M21 13v10h-21v-19h12v2h-10v15h17v-8h2zm3-12h-10.988l4.035 4-6.977 7.07 2.828 2.828 6.977-7.07 4.125 4.172v-11z"
      ></path>
    </svg>
    <ul className="github-facts">
      <li className="github-fact github-fact--version">{VERSION}</li>
      {/* <li class="github-fact github-fact--stars">1</li> */}
      {/* <li class="github-fact github-fact--forks">1</li> */}
    </ul>
  </a>
)
