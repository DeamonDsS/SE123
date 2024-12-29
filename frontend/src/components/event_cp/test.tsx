import { useState } from 'react';
import './test.css';
import { CiLogout } from "react-icons/ci";
import { FiSidebar } from "react-icons/fi";

const Sidebar = () => {
  const [isOpen, setIsOpen] = useState(true); // Sidebar initially open
  const [isContentShifted, setIsContentShifted] = useState(true); // Initially, content should shift

  const toggleSidebar = () => {
    setIsOpen(!isOpen);
    setIsContentShifted(!isContentShifted); // Shift content when sidebar is toggled
  };

  return (
    <>
      <div className={`sidebar ${isOpen ? 'open' : 'closed'}`}>
      <FiSidebar className={`sidebaricon${isOpen ? 'open' : 'closed'}`} onClick={toggleSidebar}/>
        
        <div className="sidebar-content">
          <div className="menu-section">
    
            <div className="menu-item">
              <span>Event</span>
            </div>
            <div className="menu-item">
              <span>Food and Beverage</span>
            </div>
            <div className="menu-item">
              <span>Transportation</span>
            </div>
            <div className="menu-item">
              <span>Merchandise</span>
            </div>
            <div className="menu-item">
              <span>Stock</span>
            </div>
          </div>
        </div>

        <div className="under">
          <div className="menu-item upgrade">
            <CiLogout /> 
            <span>Log Out</span>
          </div>
        </div>
      </div>
      <FiSidebar className={`sidebaricon2${isOpen ? 'open' : 'closed'}`} onClick={toggleSidebar}/>
      <div className={`content ${isContentShifted ? 'shifted' : ''}`}>
      
      </div>
      
    </>
  );
};

export default Sidebar;
