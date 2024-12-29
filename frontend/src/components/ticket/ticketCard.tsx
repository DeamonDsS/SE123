import React, { useState, useEffect } from 'react';
import './ticketCard.css';
import './button_submit.css';
import { CgErase } from "react-icons/cg";
import { Ticket } from '../../services/https/ticket';
import { getPackageById } from '../../services/https/package';
import { useFetchTickets } from '../ticketPackage/packagehook';
import { useRecoilValue } from 'recoil';
import { ticketState } from '../ticketPackage/packageinfo';
import { useFetchT } from '../ticketPackage/packagehook';
import { getTickets } from '../../services/https/ticket';

interface TicketCardProps {
  ticket: Ticket; // Define the expected prop type for the ticket
}

const TicketCard: React.FC<TicketCardProps> = ({ ticket }) => {
  const [name, setName] = useState(ticket.owner_name || ''); // Initialize name from ticket data
  const [phone, setPhone] = useState(ticket.phone || ''); // Initialize phone from ticket data
  const [packageDetails, setPackageDetails] = useState<any>(null); // State to store fetched package details
  const [ticketDetails, setTicketDetails] = useState<any>(null);
  const [loading, setLoading] = useState(true); // Loading state to handle async package fetch

  const tickets = useRecoilValue(ticketState); // Access Recoil state
  const fetchTickets = useFetchT(); // Fetch hook

  // useEffect(() => {
  //   // Use setTimeout to introduce a delay
  //   const timer = setTimeout(() => {
  //     fetchTickets(); // Trigger the fetch after the delay
  //   }, 1000); // 1000 milliseconds = 1 second delay

  //   return () => clearTimeout(timer); // Clean up the timeout on component unmount
  // }, [fetchTickets]);

  // if (!tickets.length) {
  //   return <p className="empty-state">No tickets at the moment.</p>;
  // }

  useEffect(() => {
    const fetchTickets = async () => {
      try {
        setLoading(true);
        const TicketData = await getTickets(); // Fetch the package by its ID
        setTicketDetails(TicketData); // Store the package details
      } catch (error) {
        console.error("Error fetching tickets:", error);
      } finally {
        setLoading(false);
      }
    };
    if (ticket) {
      fetchTickets(); // Fetch the package data when ticket package ID is available
    }
  }, [ticket.id]); // Dependency array ensures it triggers when `ticket.tpackage_id` changes

  useEffect(() => {
    const fetchPackage = async () => {
      try {
        setLoading(true);
        const packageData = await getPackageById(ticket.tpackage_id); // Fetch the package by its ID
        setPackageDetails(packageData); // Store the package details
      } catch (error) {
        console.error("Error fetching package:", error);
      } finally {
        setLoading(false);
      }
    };

    if (ticket.tpackage_id) {
      fetchPackage(); // Fetch the package data when ticket package ID is available
    }
  }, [ticket.tpackage_id]); // Dependency array ensures it triggers when `ticket.tpackage_id` changes

  const handleSubmit = () => {
    if (name && phone) {
      alert(`Ticket submitted for ${name} with phone: ${phone} and Order ID: ${ticket.tpackage_id}`);
    } else {
      alert('Please fill out both Name and Phone fields!');
    }
  };

  const handleClear = () => {
    setName('');
    setPhone('');
  };

  if (loading) {
    return <div>Loading package details...</div>; // Loading state message
  }

  return (
    <div className="ticket">
      <div className="ticket-left">
        <div className="left-container">
          <div className="ticket-info">
            <label>
              NAME:
              <input
                type="text"
                value={ticketDetails.owner_name}
                onChange={(e) => setName(e.target.value)}
                placeholder="Enter your name"
                className="input-field"
              />
            </label>
            <label>
              PHONE:
              <input
                type="text"
                value={ticketDetails.phone}
                onChange={(e) => setPhone(e.target.value)}
                placeholder="Enter your phone"
                className="input-field"
              />
            </label>
            <p>{ticket.order_id}</p>
          </div>
        </div>
        <div className="submiter">
          <button className="pushable" onClick={handleSubmit}>
            <span className="shadow" />
            <span className="edge" />
            <span className="front">DONE</span>
          </button>
          <button className="cancle-button" onClick={handleClear}>
            <CgErase />
          </button>
        </div>
      </div>

      <div className="ticket-right">
        <div className="ticket-header">
          <p className="club-name">Castal Kingdom</p>
          <p className="event-title">{packageDetails?.t_type} Ticket</p>
          <p className="ticketCard-price">TICKET PRICE: ฿{packageDetails?.t_price}</p>
        </div>
        <div className="prizes">
          <div className="prize">
            <p>FIRST PRIZE</p>
            <p>$100</p>
          </div>
          <div className="prize">
            <p>SECOND PRIZE</p>
            <p>$50</p>
          </div>
          <div className="prize">
            <p>THIRD PRIZE</p>
            <p>$25</p>
          </div>
        </div>
        <p className="website">www.castalkingdom.com</p>
        <p className="ticket-number">{ticket.order_id}</p>

        {/* Display package information after fetching */}
        {packageDetails && (
          <div className="package-details">
            <h3>{packageDetails.t_name}</h3>
            <p>{packageDetails.t_des}</p>
            <p>Price: ฿{packageDetails.t_price}</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default TicketCard;
