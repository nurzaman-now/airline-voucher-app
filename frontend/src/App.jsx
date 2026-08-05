import React from 'react';
import {
  Box, TextField, Typography, Paper, Container, FormControl, Select, MenuItem, InputLabel, Button, Tooltip, Grid,
  Table, TableBody, TableCell, TableContainer, TableRow
} from '@mui/material';
import AirplaneTicketIcon from '@mui/icons-material/AirplaneTicket';
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';

import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import { checkVoucher, generateVouchers } from './services/ApiClient';
import { useSnackbar } from 'notistack';


function App() {
  const { enqueueSnackbar } = useSnackbar();
  const [formResult, setFormResult] = React.useState(null);
  const [loading, setLoading] = React.useState(false);
  const [formData, setFormData] = React.useState({
    crewName: "",
    crewId: "",
    flightNumber: "",
    flightDate: null,
    aircraftType: ""
  });

  const AircraftType = {
    ATR: "ATR",
    Airbus320: "Airbus 320",
    Boeing737Max: "Boeing 737 Max"
  };

  const handleInputChange = (field) => (event) => {
    setFormData((prev) => ({
      ...prev,
      [field]: event.target.value,
    }));
  };

  const handleDateChange = (newValue) => {
    setFormData((prev) => ({
      ...prev,
      flightDate: newValue,
    }));
  };

  const handleGenerateVouchers = async (event) => {
    if (event) {
      event.preventDefault(); // Mencegah reload halaman
    }
    setLoading(true);
    try {
      // Format flightDate menjadi "DD/MM/YYYY HH:mm" agar sesuai dengan validator backend
      const formattedDate = formData.flightDate ? formData.flightDate.format('YYYY-MM-DD HH:mm:ss') : '';

      // check status voucher
      const checkStatusVoucher = await checkVoucher(formData.flightNumber, formattedDate);
      enqueueSnackbar(checkStatusVoucher.message, {
        variant: checkStatusVoucher.status === 'success' ? 'success' : 'error',
        anchorOrigin: {
          horizontal: 'center',
          vertical: 'top'
        },
        autoHideDuration: 3000,
      });
      if (checkStatusVoucher.status === 'success') {
        const dataSend = {
          ...formData,
          flightDate: formattedDate,
        }
        const generateVoucher = await generateVouchers(dataSend)
        setFormResult({
          crewName: generateVoucher?.data?.crew_name,
          crewId: generateVoucher?.data?.crew_id,
          flightNumber: generateVoucher?.data?.flight_number,
          flightDate: generateVoucher?.data?.flight_date,
          aircraftType: generateVoucher?.data?.aircraft_type,
          seat1: generateVoucher?.data?.seat1,
          seat2: generateVoucher?.data?.seat2,
          seat3: generateVoucher?.data?.seat3,
          createdAt: generateVoucher?.data?.created_at,
        })
        enqueueSnackbar(generateVoucher.message, {
          variant: generateVoucher.status === 'success' ? 'success' : 'error',
          anchorOrigin: {
            horizontal: 'center',
            vertical: 'top'
          },
          autoHideDuration: 3000,
        });
      } else {
        setFormResult({
          crewName: checkStatusVoucher?.data?.crew_name,
          crewId: checkStatusVoucher?.data?.crew_id,
          flightNumber: checkStatusVoucher?.data?.flight_number,
          flightDate: checkStatusVoucher?.data?.flight_date,
          aircraftType: checkStatusVoucher?.data?.aircraft_type,
          seat1: checkStatusVoucher?.data?.seat1,
          seat2: checkStatusVoucher?.data?.seat2,
          seat3: checkStatusVoucher?.data?.seat3,
          createdAt: checkStatusVoucher?.data?.created_at,
        })
      }

    } catch (error) {
      console.error(error)
    } finally {
      setLoading(false);
    }
  };

  return (
    <LocalizationProvider dateAdapter={AdapterDayjs}>
      <Box
        sx={{
          width: '100vw',
          minHeight: '100vh',
          backgroundImage: 'url(/airplane_bg.jpg)',
          backgroundRepeat: 'no-repeat',
          backgroundSize: 'cover',
          backgroundPosition: 'center',
          position: 'relative',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          m: 0,
          p: 0,
          '&::before': {
            content: '""',
            position: 'absolute',
            top: 0,
            left: 0,
            right: 0,
            bottom: 0,
            backgroundColor: 'rgba(0, 0, 0, 0.4)', // Overlay gelap elegan
            zIndex: 1,
          },
        }}
      >
        <Container maxWidth="md" sx={{ position: 'relative', zIndex: 2 }}>
          <Paper
            elevation={6}
            sx={{
              p: 4,
              borderRadius: 4,
              backgroundColor: 'rgba(255, 255, 255, 0.95)', // Putih bersih dengan sedikit transparansi modern
              backdropFilter: 'blur(8px)', // Efek kaca modern (glassmorphism)
              boxShadow: '0 8px 32px 0 rgba(0, 0, 0, 0.3)',
            }}
          >
            {/* Header */}
            <Box
              sx={{
                display: 'flex',
                alignItems: 'center',
                gap: 2,
                mb: 4,
                textAlign: 'left',
              }}
            >
              <img src="/logo.png" alt="Logo" style={{ height: '55px', width: 'auto', objectFit: 'contain' }} />
              <Box>
                <Typography variant="h5" fontWeight="bold" color="text.primary" sx={{ lineHeight: 1.2 }}>
                  Voucher Seat Assignment
                </Typography>
                <Typography variant="body2" color="text.secondary" sx={{ mt: 0.5 }}>
                  Mohon isi data berikut
                </Typography>
              </Box>
            </Box>

            {/* Split Grid: Kiri untuk Form, Kanan untuk Result Voucher */}
            <Grid container spacing={4}>
              {/* Kolom Kiri: Form Input */}
              <Grid item xs={12} md={6} sx={{ minWidth: '48%' }}>
                <Box
                  component="form"
                  onSubmit={handleGenerateVouchers}
                  sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}
                  autoComplete="off"
                >
                  <TextField
                    required
                    id="outlined-required-name"
                    label="Crew Name"
                    value={formData.crewName}
                    onChange={handleInputChange('crewName')}
                    fullWidth
                  />
                  <TextField
                    required
                    id="outlined-required-id"
                    label="Crew ID"
                    value={formData.crewId}
                    onChange={handleInputChange('crewId')}
                    fullWidth
                  />
                  <TextField
                    required
                    id="outlined-required-flight"
                    label="Flight Number"
                    value={formData.flightNumber}
                    onChange={handleInputChange('flightNumber')}
                    fullWidth
                  />
                  <DateTimePicker
                    label="Flight Date"
                    value={formData.flightDate}
                    onChange={handleDateChange}
                    ampm={false}
                    format="DD/MM/YYYY HH:mm"
                    slotProps={{
                      textField: {
                        required: true,
                        fullWidth: true,
                      },
                    }}
                  />
                  <FormControl fullWidth required sx={{ textAlign: 'left' }}>
                    <InputLabel id="select-aircraft-type-label" sx={{ textAlign: 'left' }}>Aircraft Type</InputLabel>
                    <Select
                      labelId="select-aircraft-type-label"
                      id="select-aircraft-type"
                      value={formData.aircraftType}
                      onChange={handleInputChange('aircraftType')}
                      label="Aircraft Type"
                      sx={{
                        textAlign: 'left',
                        '& .MuiSelect-select': {
                          textAlign: 'left',
                        },
                      }}
                    >
                      {
                        Object.values(AircraftType).map((aircraftType) => {
                          return (
                            <MenuItem key={aircraftType} value={aircraftType} sx={{ justifyContent: 'flex-start' }}>
                              {aircraftType}
                            </MenuItem>
                          );
                        })
                      }
                    </Select>
                  </FormControl>
                  <Tooltip title="Generate Voucher">
                    <Button type="submit" variant="contained" loading={loading} size="large" fullWidth>
                      <AirplaneTicketIcon sx={{ marginRight: 1 }} /> Generate
                    </Button>
                  </Tooltip>
                </Box>
              </Grid>

              {/* Kolom Kanan: Tempat/Box untuk Result Voucher */}
              <Grid item xs={12} md={6} sx={{ width: '47%' }}>
                <Box
                  sx={{
                    height: '100%',
                    minHeight: '280px',
                    border: '2px dashed #e2e8f0',
                    borderRadius: 3,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    p: 3,
                    backgroundColor: '#fafafa',
                  }}
                >
                  <Typography variant="h6" color="text.secondary" fontWeight="medium" gutterBottom>
                    Hasil Voucher
                  </Typography>
                  {
                    formResult ? (
                      <TableContainer sx={{ mt: 2 }}>
                        <Table size="small" aria-label="voucher details">
                          <TableBody>
                            <TableRow sx={{ '& td': { border: 0, py: 0.75 } }}>
                              <TableCell component="th" scope="row" sx={{ p: 0, fontWeight: 'bold', color: 'text.secondary', width: '40%' }}>
                                Crew Name
                              </TableCell>
                              <TableCell sx={{ p: 0 }}>: {formResult.crewName}</TableCell>
                            </TableRow>
                            <TableRow sx={{ '& td': { border: 0, py: 0.75 } }}>
                              <TableCell component="th" scope="row" sx={{ p: 0, fontWeight: 'bold', color: 'text.secondary' }}>
                                Crew ID
                              </TableCell>
                              <TableCell sx={{ p: 0 }}>: {formResult.crewId}</TableCell>
                            </TableRow>
                            <TableRow sx={{ '& td': { border: 0, py: 0.75 } }}>
                              <TableCell component="th" scope="row" sx={{ p: 0, fontWeight: 'bold', color: 'text.secondary' }}>
                                Flight Number
                              </TableCell>
                              <TableCell sx={{ p: 0 }}>: {formResult.flightNumber}</TableCell>
                            </TableRow>
                            <TableRow sx={{ '& td': { border: 0, py: 0.75 } }}>
                              <TableCell component="th" scope="row" sx={{ p: 0, fontWeight: 'bold', color: 'text.secondary' }}>
                                Flight Date
                              </TableCell>
                              <TableCell sx={{ p: 0 }}>: {formResult.flightDate}</TableCell>
                            </TableRow>
                            <TableRow sx={{ '& td': { border: 0, py: 0.75 } }}>
                              <TableCell component="th" scope="row" sx={{ p: 0, fontWeight: 'bold', color: 'text.secondary' }}>
                                Aircraft Type
                              </TableCell>
                              <TableCell sx={{ p: 0 }}>: {formResult.aircraftType}</TableCell>
                            </TableRow>
                            <TableRow sx={{ '& td': { border: 0, py: 0.75 } }}>
                              <TableCell sx={{ p: 0 }} colSpan={2}>
                                <Box sx={{ display: 'flex', gap: 1, mt: 0.5, flexDirection: 'column', flexWrap: 'wrap' }}>
                                  {[formResult.seat1, formResult.seat2, formResult.seat3].map((seat, index) => (
                                    <Box
                                      key={index}
                                      sx={{
                                        display: 'inline-flex',
                                        alignItems: 'center',
                                        gap: 0.5,
                                        backgroundColor: '#eff6ff',
                                        color: '#1d4ed8',
                                        border: '1px solid #bfdbfe',
                                        borderRadius: 2,
                                        px: 1.5,
                                        py: 0.5,
                                        fontWeight: 'bold',
                                        fontSize: '0.85rem',
                                        boxShadow: '0 2px 4px rgba(29, 78, 216, 0.05)',
                                      }}
                                    >
                                      Seat {index + 1}: {seat}
                                    </Box>
                                  ))}
                                </Box>
                              </TableCell>
                            </TableRow>
                          </TableBody>
                        </Table>
                      </TableContainer>
                    ) : (
                      <Box sx={{ flexGrow: 1, display: 'flex', alignItems: 'center', justifyContent: 'center', mt: 4 }}>
                        <Typography variant="body2" color="text.disabled" align="center">
                          Detail voucher dan alokasi kursi penerbangan akan muncul di sini setelah Anda menekan tombol Generate.
                        </Typography>
                      </Box>
                    )
                  }
                </Box>
              </Grid>
            </Grid>
          </Paper>
        </Container>
      </Box>
    </LocalizationProvider>
  );
}

export default App;
