import './ticket_main.css';
import TicketPackage from '../../components/ticketPackage/tpackage';
import { AiFillCaretLeft } from "react-icons/ai";

const TicketMain = () => {
  const handleReturn = () => {
    // Logic for return action (e.g., navigating to a previous page)
    window.history.back();
  };

  return (
    <>
      <div className="base">
        <div className="top">
          <button className="t-return" onClick={handleReturn}>
            <AiFillCaretLeft size={24} color="#ffffff" />
          </button>
          {/* <h1>Test</h1> */}
        </div>
        <div className="tcontainer">
          <TicketPackage />
        </div>
      </div>
    </>
  );
};

export default TicketMain;
