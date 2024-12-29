import { atom } from "recoil";
import { Tpackage } from "../../services/https/package";
import { Ticket } from "../../services/https/ticket";

export const packageState = atom<Tpackage[]>({
  key: "packageState", // Unique key for the atom
  default: [],         // Initial empty state
});

export const ticketState = atom<Ticket[]>({
  key: "ticketState",
  default: [],
});
