import './countTicket.css';

interface CountTicketProps {
  type: string; // Ticket type (e.g., Basic, Standard, Premium, VIP)
  count: number; // Number of tickets
}

const CountTicket: React.FC<CountTicketProps> = ({ type, count }) => {
  // Assign an icon based on the ticket type
  const getIcon = (ticketType: string) => {
    switch (ticketType) {
      case 'Basic':
        return '🥞';
      case 'Standard':
        return '🎟️';
      case 'Premium':
        return '✨';
      case 'VIP':
        return '👑';
      default:
        return '🎫';
    }
  };

  return (
    <div className="count-base">
      {/* Header Section */}
      <div className="count-head">
        <div className="ticket-icon">{getIcon(type)}</div>
        <div className="ticket-title">{type} Ticket</div>
      </div>

      {/* Body Section */}
      <div className="count-body">
        <span>x</span>
        <div className="number-ticket">{count}</div>
      </div>
    </div>
  );
};

export default CountTicket;
