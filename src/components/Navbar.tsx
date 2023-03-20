import React, { useState } from 'react'
import { FaGithub } from "react-icons/fa"
import { FiBook } from "react-icons/fi"

const Navbar = () => {

  return (
    <header>
      <nav>
        <ul className="flex justify-between items-center">
          <li><a href="#" className='flex items-center'><FiBook size="2.5rem" /> Dictionary</a></li>
          <li><a href="https://github.com/jjalbuenacabuyao/dictionary" target='_blank'><FaGithub size="1.5rem" /></a></li>
        </ul>
      </nav>
    </header>
  )
}

export default Navbar