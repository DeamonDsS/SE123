import React, { useState } from 'react';
import './order_main.css';
import TicketCard from '../../components/ticket/ticketCard';
import Checkout from '../../components/ticket/checkout';
import CountTicket from '../../components/ticket/countTicket';
const TICKET_PACKAGES = ['Basic', 'Standard', 'Premium', 'VIP']; // Fixed 4 ticket packages

const OrderMain = () => {
  // State for managing tickets and their quantities
  const [tickets, setTickets] = useState<any[]>([]);
  const [ticketCounts, setTicketCounts] = useState<{ [key: string]: number }>({
    Basic: 0,
    Standard: 0,
    Premium: 0,
    VIP: 0,
  });

  // Function to add a new ticket of a specific type
  const addTicket = (type: string) => {
    setTickets((prevTickets) => [
      ...prevTickets,
      {
        id: Date.now(),
        type,
        name: '',
        phone: '',
        orderId: `NO. ${Math.floor(100000 + Math.random() * 900000)}`,
      },
    ]);

    // Update ticket count
    setTicketCounts((prevCounts) => ({
      ...prevCounts,
      [type]: prevCounts[type] + 1,
    }));
  };

  return (
    <div className="order-base">
      <div className="order-top">
        <p>Order Page</p>
        {/* Buttons to Add Tickets for Each Package */}
        <div className="ticket-buttons">
          {TICKET_PACKAGES.map((type) => (
            <button
              key={type}
              className="add-ticket-button"
              onClick={() => addTicket(type)}
            >
              Add {type} Ticket
            </button>
          ))}
        </div>
      </div>

      <div className="order-container">
        {/* Left Box for Checkout */}
        <div className="left-box">
          <Checkout />
        </div>

        {/* Right Box */}
        <div className="right-box">
          <div className="tickets-box">
            {/* Render Tickets Dynamically */}
            {tickets.map((ticket) => (
              <TicketCard key={ticket.id} ticket={ticket} />
            ))}
          </div>

          {/* Fixed 4 CountTicket Boxes */}
          <div className="info-box">
            <div className="info-left">
              {TICKET_PACKAGES.map((type) => (
                <CountTicket key={type} type={type} count={ticketCounts[type]} />
              ))}
            </div>
            <div className="info-right"></div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default OrderMain;
