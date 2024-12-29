import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; // Update with your actual API base URL

// Retrieve token and token type from localStorage
const getAuthHeaders = () => {
  const token = localStorage.getItem("token");
  const tokenType = localStorage.getItem("token_type");
  return {
    "Content-Type": "application/json",
    Authorization: `${tokenType} ${token}`,
  };
};

// Retrieve all orders
export const getAllOrders = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/orders`, {
      headers: getAuthHeaders(),
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching orders:', error);
    throw error;
  }
};

// Retrieve a specific order by ID
export const getOrderById = async (id: number) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/orders/${id}`, {
      headers: getAuthHeaders(),
    });
    return response.data;
  } catch (error) {
    console.error(`Error fetching order with ID ${id}:`, error);
    throw error;
  }
};

// Create a new order
export const createOrder = async (orderData: any) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/orders`, orderData, {
      headers: getAuthHeaders(),
    });
    return response.data;
  } catch (error) {
    console.error('Error creating order:', error);
    throw error;
  }
};

// Update an existing order
export const updateOrder = async (id: number, orderData: any) => {
  try {
    const response = await axios.put(`${API_BASE_URL}/orders/${id}`, orderData, {
      headers: getAuthHeaders(),
    });
    return response.data;
  } catch (error) {
    console.error(`Error updating order with ID ${id}:`, error);
    throw error;
  }
};

// Delete an order by ID
export const deleteOrder = async (id: number) => {
  try {
    const response = await axios.delete(`${API_BASE_URL}/orders/${id}`, {
      headers: getAuthHeaders(),
    });
    return response.data;
  } catch (error) {
    console.error(`Error deleting order with ID ${id}:`, error);
    throw error;
  }
};
