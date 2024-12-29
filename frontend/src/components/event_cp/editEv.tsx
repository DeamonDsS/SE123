import React, { useState } from 'react';
import { createEvent } from '../../services/https/event_sv/event_sv';
import './editEv.css';

const EditForm: React.FC = () => {
  const [formData, setFormData] = useState({
    EventName: '',
    location: '',
    detail: '',
    startDate: '',
    endDate: '',
    status: 'private',
    image: null as File | null, // Add image file state
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleStatusChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData((prevData) => ({
      ...prevData,
      status: e.target.value,
    }));
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      setFormData((prevData) => ({
        ...prevData,
        image: e.target.files[0], // Save the first selected file
      }));
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Validation: ensure required fields are filled
    if (!formData.EventName || !formData.location || !formData.startDate) {
      alert('กรุณากรอกข้อมูลให้ครบถ้วน');
      return;
    }

    // Prepare FormData for submission, including the image
    const eventData = new FormData();
    eventData.append('event_name', formData.EventName);
    eventData.append('location_id', formData.location);
    eventData.append('detail', formData.detail);
    eventData.append('start_event', formData.startDate);
    eventData.append('end_event', formData.endDate);
    eventData.append('ispublic', formData.status);
 
    if (formData.image) {
      eventData.append('cover', formData.image);
    }

    try {
      const response = await createEvent(eventData);
      console.log('Event created:', response);
    } catch (error) {
      console.error('Error submitting form:', error);
    }
  };

  return (
    <div className="edit-container">
      <button className="back-button">◀</button>
      <form className="edit-form" onSubmit={handleSubmit}>
        <div className="form-group">
          <label>ชื่องาน</label>
          <input
            type="text"
            name="EventName"
            value={formData.EventName}
            onChange={handleChange}
          />
        </div>

        <div className="form-group">
          <label>รูปปก</label>
          <input type="file" onChange={handleFileChange} />
        </div>

        <div className="form-group">
          <label>สถานที่</label>
          <input
            type="text"
            name="location"
            value={formData.location}
            onChange={handleChange}
          />
        </div>

        <div className="form-group">
          <label>รายละเอียด</label>
          <textarea
            name="description"
            value={formData.detail}
            onChange={handleChange}
          />
        </div>

        <div className="form-group">
          <label>ระยะเวลาจัดงาน</label>
          <div className="date-range">
            <input
              type="date"
              name="startDate"
              value={formData.startDate}
              onChange={handleChange}
            />
            <span> - </span>
            <input
              type="date"
              name="endDate"
              value={formData.endDate}
              onChange={handleChange}
            />
          </div>
        </div>

        <div className="form-group">
          <label>สถานะ</label>
          <div className="status-options">
            <label>
              <input
                type="radio"
                name="status"
                value="private"
                checked={formData.status === 'private'}
                onChange={handleStatusChange}
              />
              ซ่อน
            </label>
            <label>
              <input
                type="radio"
                name="status"
                value="public"
                checked={formData.status === 'public'}
                onChange={handleStatusChange}
              />
              เผยแพร่
            </label>
          </div>
        </div>

        <button type="submit" className="submit-button">บันทึก</button>
      </form>
    </div>
  );
};

export default EditForm;


