import axios from "axios";

const BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/api";

const apiClient = axios.create({
  baseURL: BASE_URL,
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
  // Izinkan axios menyelesaikan request (tidak melempar error) untuk semua HTTP status code di bawah 500
  validateStatus: (status) => status < 500,
});



export const checkVoucher = async (
  flightNumber,
  flightDate,
) => {
  try {
    // Menggunakan variabel baseUrl untuk request Axios
    const response = await apiClient.post(`/check`, {
      flight_number: flightNumber,
      flight_date: flightDate,
    });
    return response.data;
  } catch (error) {
    return error;
  }
};

export const generateVouchers = async ({
  crewName,
  crewId,
  flightNumber,
  flightDate,
  aircraftType
}) => {
  try {
    const response = await apiClient.post(`/generate`, {
      crew_name: crewName,
      crew_id: crewId,
      flight_number: flightNumber,
      flight_date: flightDate,
      aircraft_type: aircraftType,
    });
    return response.data;
  } catch (error) {
    return error;
  }
};
