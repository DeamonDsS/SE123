import './event.css'
import EvetBox from '../../../components/event_cp/evetBox'
import Sidebar from '../../../components/event_cp/test'
// import EditForm from '../../../components/event_cp/editEv'
const Event = () => {
  return (
    <div className="wrapper">
       
            <Sidebar />
            <EvetBox/>
            {/* <EditForm /> */}

        
     
    </div>
  )
}

export default Event