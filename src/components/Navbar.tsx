import React, { useState } from 'react'
import { FaGithub } from "react-icons/fa"
import { FiBook } from "react-icons/fi"
import { MdLightMode, MdDarkMode } from "react-icons/md"


const Navbar = () => {
  const htmlClassList : DOMTokenList = document.documentElement.classList;
  const [theme, setTheme] = useState("");
  let isDark : boolean = theme === "dark";

  isDark ? htmlClassList.add("dark")
         : htmlClassList.remove("dark");

  return (
    <header className="flex items-center justify-between">
      <a href="#" className='flex items-center'><FiBook size="2.5rem" /> Dictionary</a>
      <nav>
        <ul className="flex justify-between items-center">
          <li className='grid place-items-center'>
            <button onClick={() => setTheme(isDark ? "light" : "dark")}>
              { isDark ? <MdLightMode size="1.5rem" /> : <MdDarkMode size="1.5rem" /> }
            </button>
          </li>
          <li>
            <a href="https://github.com/jjalbuenacabuyao/dictionary" target='_blank'><FaGithub size="1.5rem" /></a>
          </li>
        </ul>
      </nav>
    </header>
  )
}

export default Navbar