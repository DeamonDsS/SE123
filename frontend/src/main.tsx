import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { RecoilRoot } from 'recoil';
import Event from './pages/admin/event/event';
import SignInPages from './pages/authentication/Login';
import TicketMain from './pages/ticket_page/ticket_main';
import OrderMain from './pages/order_page/order_main';
import "./index.css";


const router = createBrowserRouter([
  
  {path: "/", element: <Event />,},
  {path: "/singin", element: <SignInPages />,},
  {path: "/ticketpage", element: <TicketMain />,},
  {path: "/orderpage", element: <OrderMain />,},


    
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RecoilRoot>
        <RouterProvider router={router} />
    </RecoilRoot>
  </React.StrictMode>
);








