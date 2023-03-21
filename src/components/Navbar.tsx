import React, { useState, useEffect } from 'react'
import { FaGithub } from "react-icons/fa"
import { FiBook } from "react-icons/fi"
import { MdLightMode, MdDarkMode } from "react-icons/md"


const Navbar = () => {
  const [theme, setTheme] = useState("dark");
  let isDark : boolean = theme === "dark";

  useEffect(() => {
    const htmlClassList : DOMTokenList = document.documentElement.classList;
    isDark ? htmlClassList.add("dark")
           : htmlClassList.remove("dark");
  }, [theme])

  return (
    <header className="flex items-center justify-between px-6 py-5">
      <a href="#" className='flex items-center gap-2'><FiBook size="2.5rem" /> Dictionary</a>
      <nav>
        <ul className="flex justify-between items-center gap-2">
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