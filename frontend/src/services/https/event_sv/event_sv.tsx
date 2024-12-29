import axios from "axios";

// URL สำหรับ API
const apiUrl = "http://localhost:8000";

// ดึง Token และ Token Type จาก localStorage
const Authorization = localStorage.getItem("token");
const Bearer = localStorage.getItem("token_type");

// ตัวเลือกในการตั้งค่า headers
const requestOptions = {
  headers: {
    "Content-Type": "application/json",
    Authorization: `${Bearer} ${Authorization}`, // รวม Token และ Token Type
  },
};

export interface Event {
    event_name: string;
    detail: string;
    cover: string;
    ispublic: number;
    start_event: string;
    end_event: string;
    type_id: number;
    location_id: number;
    admin_id: number;
  }

// ฟังก์ชันดึงข้อมูลทั้งหมดของ events (GET /events)
export async function getEvents() {
  try {
    const response = await axios.get(`${apiUrl}/events`, requestOptions);
    return response.data; // ส่งคืนข้อมูลทั้งหมดของ events
  } catch (error) {
    console.error("Error fetching events:", error);
    throw error;
  }
}

// ฟังก์ชันดึงข้อมูลของ event ตาม id (GET /events/:id)
export async function getEvent(id: string) {
  try {
    const response = await axios.get(`${apiUrl}/events/${id}`, requestOptions);
    return response.data; // ส่งคืนข้อมูลของ event ตาม id
  } catch (error) {
    console.error("Error fetching event:", error);
    throw error;
  }
}

// ฟังก์ชันสร้าง event ใหม่ (POST /events)
export async function createEvent(data: Event) {
  try {
    const response = await axios.post(`${apiUrl}/events`, data, requestOptions);
    return response.data; // ส่งคืนผลลัพธ์จากการสร้าง event ใหม่
  } catch (error) {
    console.error("Error creating event:", error);
    throw error;
  }
}

// ฟังก์ชันอัพเดต event ตาม id (PUT /events/:id)
export async function updateEvent(id: string, data: Event) {
  try {
    const response = await axios.put(`${apiUrl}/events/${id}`, data, requestOptions);
    return response.data; // ส่งคืนผลลัพธ์จากการอัพเดต event
  } catch (error) {
    console.error("Error updating event:", error);
    throw error;
  }
}

// ฟังก์ชันลบ event ตาม id (DELETE /events/:id)
export async function deleteEvent(id: string) {
  try {
    const response = await axios.delete(`${apiUrl}/events/${id}`, requestOptions);
    return response.data; // ส่งคืนผลลัพธ์จากการลบ event
  } catch (error) {
    console.error("Error deleting event:", error);
    throw error;
  }
}
