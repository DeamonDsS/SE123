import axios from "axios";

// Base API URL
const apiUrl = "http://localhost:8000";

// Helper for setting headers with token-based authentication
const token = localStorage.getItem("token");
const tokenType = localStorage.getItem("token_type");

const requestOptions = {
  headers: {
    "Content-Type": "application/json",
    Authorization: `${tokenType} ${token}`,
  },
};

// Interface for Tpackage
export interface Tpackage {
  ID: number; // Optional for creation
  t_name: string;
  t_type: string;
  t_price: number;
  t_zone: string;
  t_des: string;
}

// Get all packages
export async function getAllPackages(): Promise<Tpackage[]> {
  try {
    const response = await axios.get<Tpackage[]>(`${apiUrl}/tpackages`, requestOptions);
    return response.data; // `response.data` is now typed as `Tpackage[]`
  } catch (error) {
    console.error("Error fetching packages:", error);
    throw error;
  }
}

// Get a single package by ID
export async function getPackageById(id: number | string): Promise<Tpackage> {
  try {
    const response = await axios.get<Tpackage>(`${apiUrl}/tpackages/${id}`, requestOptions);
    return response.data; // `response.data` is now typed as `Tpackage`
  } catch (error) {
    console.error("Error fetching package by ID:", error);
    throw error;
  }
}

// Create a new package
export async function createPackage(data: Tpackage): Promise<Tpackage> {
  try {
    const response = await axios.post<Tpackage>(`${apiUrl}/tpackages`, data, requestOptions);
    return response.data; // `response.data` is now typed as `Tpackage`
  } catch (error) {
    console.error("Error creating package:", error);
    throw error;
  }
}

// Update an existing package
export async function updatePackage(
  id: number | string,
  data: Tpackage
): Promise<Tpackage> {
  try {
    const response = await axios.put<Tpackage>(
      `${apiUrl}/tpackages/${id}`,
      data,
      requestOptions
    );
    return response.data; // `response.data` is now typed as `Tpackage`
  } catch (error) {
    console.error("Error updating package:", error);
    throw error;
  }
}

// Delete a package by ID
export async function deletePackage(id: number | string): Promise<void> {
  try {
    await axios.delete(`${apiUrl}/tpackages/${id}`, requestOptions);
  } catch (error) {
    console.error("Error deleting package:", error);
    throw error;
  }
}
