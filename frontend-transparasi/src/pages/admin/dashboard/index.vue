<template>
  <v-container >

    <div class=" ">
        <div class="d-flex">
          <div class="chart-line-dashboard" style="">
            <div class="d-flex justify-space-between">
              <h3 class="pr-5">Statistik Keluhan</h3>
              <div class="d-flex">
                <v-btn class="mr-1" color="#BF3232">Bulanan</v-btn>
                <v-btn class="" variant="outlined" style="color: black; border-color: #BF3232 ;">Tahunan</v-btn>
              </div>
            </div>
            <v-divider class="mt-3"></v-divider>
            <apexchart class="mt-3" width="100%" height="90%" type="line" :options="options" :series="series"></apexchart>
          </div>
        </div>
    </div>

    <div class="mt-8">
      <v-row>
        <v-col md="12" lg="4">
          <v-card variant="outlined" style="border-color: #BF3232;">
            <div class="d-flex justify-space-between dashboard-card" style="color: black">
              <div class="title">
                <h4>Total Keluhan</h4>
              </div>
              <div class="data-angka">
                <h1>100</h1>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col md="12" lg="4">
          <v-card style="background-color: #387144">
            <div class="d-flex justify-space-between dashboard-card">
              <div class="title">
                <h4>Keluhan Sudah Ditanggapi</h4>
              </div>
              <div class="data-angka">
                <h1>100</h1>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col md="12" lg="4">
          <v-card style="background-color: #BFAD32;">
            <div class="d-flex justify-space-between dashboard-card">
              <div class="title">
                <h4>Keluhan Belum Ditanggapi</h4>
              </div>
              <div class="data-angka">
                <h1>100</h1>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>
    </div>

    <div class="mt-3">
        <div class="d-flex align-center justify-space-between mr-3 mt-2">
          <v-card-title><b>Aktifitas Terbaru</b></v-card-title>
          <a href="">
            <v-icon class="">mdi-dots-vertical</v-icon>
          </a>
        </div>
        <v-divider class="mx-2"></v-divider>
        <div class="">
          <v-table>
            <thead>
              <tr class="">
                <th class="text-left font-weight-bold">
                  Tanggal
                </th>
                <th class="text-left font-weight-bold">
                  Pengguna
                </th>
                <th class="text-left font-weight-bold">
                  Aktivitas
                </th>
                <th class="text-left font-weight-bold">
                  Status
                </th>
              </tr>
            </thead>
              <tbody>
                <tr
                v-for="item in log"
                :key="item.name"
                >
              <td>{{ item.CreatedAt }}</td>
              <td>{{ item.User.username }}</td>
              <td>{{ item.aktivitas }}</td>
              <td>{{ item.status }}</td>

              </tr>
              </tbody>
          </v-table>
        </div>
    </div>

  </v-container>
</template>
<script>
import cardDashboard from '@/components/card-dashboard.vue';
import axios from 'axios';

export default{
  components: { cardDashboard },
  data() {
    return {
      log:[],
      options: {
        chart: {
          id: 'vuechart-example',
          zoom: {
                enabled: false
              }
        },
        grid: {
              row: {
                colors: ['#FFE3E3', 'transparent'], // takes an array which will be repeated on columns
                opacity: 0.5
              },
            },
        xaxis: {
          categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998]
        },
        yaxis: {
          min: 0, // Nilai minimum di sumbu Y
          max: 150, // Nilai maksimum di sumbu Y
          tickAmount: 3, // Membagi skala dengan kelipatan 50
          labels: {
            formatter: function (value) {
              return value.toFixed(0); // Tampilkan angka bulat
            }
          },
        },
      },
      series: [{
        name: 'series-1',
        data: [30, 35, 45, 50, 49, 80, 70, 150]
      }],
      today : this.todayDate()
    }
  },
  mounted(){
    this.historyLog();
  },
  methods: {
    todayDate(){
      const date = new Date();
      const months = [
        'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
        'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'
    ];
    const year = date.getFullYear();
    const month = months[date.getMonth()]; // Mendapatkan nama bulan singkat
    const day = String(date.getDate()).padStart(2, '0'); // Mendapatkan tanggal dan memastikan dua digit
    return `${month}, ${day} ${year}`; // Format: "Oct, 10 2024"
    },
    historyLog(){
      axios.get('/api/log')
      .then(res=>{
        console.log(res.data.data)
        this.log = res.data.data
      })
    }
  },
  
}
</script>

<style scooped>

</style>
