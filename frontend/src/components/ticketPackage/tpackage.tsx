import './tpackage.css';
import { useEffect } from 'react';
import { useRecoilValue } from 'recoil';
import { packageState } from './packageinfo';
import { useFetchPackages } from './packagehook';
import { createTicket } from '../../services/https/ticket';



const TicketPackage: React.FC = () => {
  const packages = useRecoilValue(packageState); // Access Recoil state
  // const [isLoading, setIsLoading] = useState(true);
  const fetchPackages = useFetchPackages(); // Fetch hook

  useEffect(() => {
    // Use setTimeout to introduce a delay
    const timer = setTimeout(() => {
      fetchPackages(); // Trigger the fetch after the delay
    }, 1000); // 1000 milliseconds = 1 second delay

    return () => clearTimeout(timer); // Clean up the timeout on component unmount
  }, [fetchPackages]);

  if (!packages.length) {
    return <p className="empty-state">No packages available at the moment.</p>;
  }

  const handleBuy = async (ID: number) => {
    try {
      alert(`You selected the ${ID} package!`);
      const ticketData = {
        owner_name: "", // Replace with actual user data
        phone:"",
        code_id: 1, // Replace with actual ticket code logic
        tpackage_id: ID,
        order_id: 1 // Replace with actual order id
      };
      
      const newTicket = await createTicket(ticketData);
      alert('Ticket created successfully!'); // Notify user
      // Handle additional UI updates here, e.g., redirect to order page, update cart UI, etc.
    } catch (error) {
      console.error("Error creating ticket:", error);
      alert('Failed to create ticket.');
    }
  };

  return (
    <div className="ticket-container">
      {packages.map((pkg, index) => (
        <div className="ticket-box" key={index}>
          <div className="ticket-header">
            {pkg.t_name}
          </div>
          <div className="ticket-price">
            <b className='ticket-price-style'>฿{pkg.t_price}/Ticket</b>
          </div>
          <div className="description-container">
            <div className="zone-description">- {pkg.t_zone}</div>
            <div className="ticket-description">
              <p>{pkg.t_des}</p>
            </div>
          </div>
          <button
            className="ticket-button"
            aria-label={`Buy ${pkg.t_name}`}
            onClick={() => handleBuy(pkg.ID)}
          >
            Add to Cart
          </button>
        </div>
      ))}
    </div>
  );
};

export default TicketPackage;
