import React, { FC } from 'react'
import { VERSION } from '@site/src/constants/githubFacts'

export const GithubLinkMob: FC = () => (
  <>
    <a
      className="menu__link header-github-link header-github-link-mob"
      href="https://github.com/H-BF/sgroups"
      target="_blank"
      rel="noopener noreferrer"
      aria-label="GitHub repository"
    >
      H-BF/sgroups
    </a>
    <ul className="github-facts github-facts-mob">
      <li className="github-fact github-fact--version">{VERSION}</li>
      {/* <li class="github-fact github-fact--stars">1</li> */}
      {/* <li class="github-fact github-fact--forks">1</li> */}
    </ul>
  </>
)
