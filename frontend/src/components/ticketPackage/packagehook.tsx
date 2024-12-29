import { useSetRecoilState } from "recoil";
import { packageState } from "./packageinfo";
import { ticketState } from "./packageinfo";
import { getAllPackages } from "../../services/https/package";
import { getTicketsL, Ticket } from "../../services/https/ticket";

export function useFetchPackages() {
  const setPackages = useSetRecoilState(packageState);

  const fetchPackages = async () => {
    try {
      const packages = await getAllPackages(); // Fetch packages from the API
      setPackages(packages); // Update the Recoil atom
    } catch (error) {
      console.error("Failed to fetch packages:", error);
    }
  };

  return fetchPackages;
}

export function useFetchT() {
  const setT = useSetRecoilState(ticketState);

  const fetchPackages = async () => {
    try {
      const tickets = await getTicketsL(); // Fetch packages from the API
      setT(tickets); // Update the Recoil atom
    } catch (error) {
      console.error("Failed to fetch packages:", error);
    }
  };

  return fetchPackages;
}

export function useFetchTickets() {
  const setTickets = useSetRecoilState(ticketState);

  const fetchTickets = async (): Promise<Ticket[]> => {
    try {
      const tickets = await getTicketsL(); // Fetch tickets from the API
      setTickets(tickets); // Update Recoil state with fetched tickets
      return tickets; // Return the fetched tickets
    } catch (error) {
      console.error('Failed to fetch tickets:', error);
      return []; // Return an empty array on error
    }
  };

  return fetchTickets;
}