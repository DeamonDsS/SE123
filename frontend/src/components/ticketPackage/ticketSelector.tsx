import { selector } from "recoil";
import { ticketState } from "./packageinfo"; // Path to your atom definition

export const ticketsByOrderIdSelector = selector({
  key: "ticketsByOrderIdSelector",
  get: ({ get }) => {
    const tickets = get(ticketState);
    return (orderId: number) => tickets.filter(ticket => ticket.order_id === orderId);
  },
});
