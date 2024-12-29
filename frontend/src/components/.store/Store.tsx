import { atom } from "recoil";
interface User{
    name?:string,
    detail?:number
}
interface EventData{
    nameEvent?:string,
    start?:string,
    stop?:string,
}
export const userInfo = atom<User>({
    key:'user',
    default:{
        name:'Best',
        detail:5
    
    }
})


export const EventData = atom({
  key: 'event', // Unique key for this atom
  default: [
    { nameEvent: 'งานบวช', start: '10/12/2024', stop: '10/12/2024' },
    { nameEvent: 'งานวันเกิดรถไฟ', start: '10/12/2024', stop: '10/12/2024' },
    { nameEvent: 'งานกินต้นไม้', start: '10/12/2024', stop: '10/12/2024' },
    { nameEvent: 'งานแต่งงาน', start: '15/12/2024', stop: '15/12/2024' },
    { nameEvent: 'งานปีใหม่', start: '31/12/2024', stop: '01/01/2025' },
    { nameEvent: 'งานประเพณีสงกรานต์', start: '13/04/2025', stop: '15/04/2025' },
    { nameEvent: 'งานประจำปีเทศกาลโคมไฟ', start: '28/11/2024', stop: '30/11/2024' },
    { nameEvent: 'งานกีฬา', start: '05/11/2024', stop: '07/11/2024' },
    { nameEvent: 'งานดนตรีกลางแจ้ง', start: '20/11/2024', stop: '21/11/2024' },
    { nameEvent: 'งานเลี้ยงบริษัท', start: '22/12/2024', stop: '22/12/2024' },
    { nameEvent: 'งานแสดงสินค้า', start: '10/11/2024', stop: '12/11/2024' },
    { nameEvent: 'งานครบรอบบริษัท', start: '30/12/2024', stop: '30/12/2024' },
    { nameEvent: 'งานประชุมใหญ่', start: '03/01/2025', stop: '04/01/2025' },
    { nameEvent: 'งานแฟชั่นโชว์', start: '15/11/2024', stop: '16/11/2024' },
    { nameEvent: 'งานแสดงศิลปะ', start: '01/12/2024', stop: '05/12/2024' },
    { nameEvent: 'งานหมั้น', start: '18/12/2024', stop: '18/12/2024' },
    { nameEvent: 'งานสัมมนา', start: '12/12/2024', stop: '13/12/2024' },
    { nameEvent: 'งานรวมญาติ', start: '25/12/2024', stop: '25/12/2024' },
 
  ]
});
