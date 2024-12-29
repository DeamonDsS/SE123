import { useRecoilValue } from 'recoil';
import { EventData } from '../.store/Store';  
import './styleEvent.css';
import Cat from '../../../public/cat.jpg';
import { useState } from 'react';
import { FaLessThan,FaGreaterThan  } from "react-icons/fa6";

const CradEvent = () => {
  const events = useRecoilValue(EventData); // Fetch event data from Recoil
  const [currentPage, setCurrentPage] = useState(1); // Track the current page
  const eventsPerPage = 5; // 5 events per page

  // Calculate the index range for slicing the events array
  const indexOfLastEvent = currentPage * eventsPerPage; 
  const indexOfFirstEvent = indexOfLastEvent - eventsPerPage; 
  const currentEvents = events.slice(indexOfFirstEvent, indexOfLastEvent); // Slice events for the current page

  // Calculate total number of pages
  const totalPages = Math.ceil(events.length / eventsPerPage);

  // Handle page change (when user clicks next or previous)
  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  return (
    <div >
      {currentEvents.map((event, index) => (
        <div key={index} className="wrapperCreadEvant">
        <div className="headBoxEv">
          <img id="picCardEveat" src={Cat} alt="Event" />
          <div className="inheadBoxEv">
            <b>
              <div>{event.nameEvent}</div>
            </b>
            <div>
              {event.start} - {event.stop}
            </div>
          </div>
           </div>
          <div className="tailBoxEv">
            <button className="editEVbt">Edit</button>
            <button className="deleteEVbt">Delete</button>
          </div>
       
        </div>
      ))}

      {/* Pagination Controls */}
      <div className="pagination">
        <button
          onClick={() => handlePageChange(currentPage - 1)}
          disabled={currentPage === 1}
        >
        <FaLessThan/>
        </button>
        {[...Array(totalPages)].map((_, index) => (
          <button
            key={index}
            onClick={() => handlePageChange(index + 1)}
            className={currentPage === index + 1 ? 'active' : ''}
          >
            {index + 1}
          </button>
        ))}
        <button
          onClick={() => handlePageChange(currentPage + 1)}
          disabled={currentPage === totalPages}
        >
          <FaGreaterThan />
        </button>
      </div>
    </div>
  );
};

export default CradEvent;