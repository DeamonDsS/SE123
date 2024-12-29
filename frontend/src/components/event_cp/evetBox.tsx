import CradEvent from "./cradEvent"
import './styleEvent.css'
import { IoIosAddCircleOutline } from "react-icons/io";
const EvetBox = () => {
  return (
    <>
      <div className="Event"><h2>Event</h2><IoIosAddCircleOutline className="addicon"/></div>
      <hr/>
     <div className="evetbox">
   
        <div><CradEvent /></div>
        
    </div>
    </>
   
    
  )
}

export default EvetBox