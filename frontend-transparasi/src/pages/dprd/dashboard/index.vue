<template>
  <v-container >
  <div class="pa-5">
    <v-card>
    <div class="pa-5">
        <div class="d-flex">
          <div class="chart-line-dashboard" style="">
            <div class="d-flex justify-space-between">
              <h3 class="pr-5">Statistik Keluhan</h3>
              <div class="d-flex">
                <v-btn class="mr-1" color="#BF3232" @click="getAduanStatsBulan()">Bulanan</v-btn>
                <v-btn class="" variant="outlined" style="color: black; border-color: #BF3232 ;" @click="getAduanStatsTahun()">Tahunan</v-btn>
              </div>
            </div>
            <v-divider class="mt-3"></v-divider>
            <apexchart class="mt-3" width="100%" height="90%" type="line" :options="options" :series="series"></apexchart>
          </div>
        </div>
    </div>
  </v-card>

    <div class="mt-8 mb-8">
      <v-row>
        <v-col col="12" lg="4">
          <v-card class="text-white" style="background-color: #FF4242; height: 100%;">
            <div class="d-flex justify-space-between dashboard-card">
              <div class="title">
                <h4>Total Keluhan</h4>
              </div>
              <div class="data-angka">
                <h1>{{ counterAduan.total }}</h1>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col md="6" lg="4">
          <v-card style="background-color: #4A975B">
            <div class="d-flex justify-space-between dashboard-card text-white">
              <div class="title">
                <h4>Keluhan Sudah Ditanggapi</h4>
              </div>
              <div class="data-angka">
                <h1>{{ counterAduan.sudah_ditanggapi }}</h1>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col md="6" lg="4">
          <v-card style="background-color: #FFE642;">
            <div class="d-flex justify-space-between dashboard-card text-white">
              <div class="title">
                <h4>Keluhan Belum Ditanggapi</h4>
              </div>
              <div class="data-angka">
                <h1>{{ counterAduan.belum_ditanggapi }}</h1>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>
    </div>

    <div class="mt-3">
      <v-card
      elevation="4"
      style="background-color: #FAFAFA;"
      >
        <div class="d-flex align-center justify-start mr-3 mt-2">
          <v-card-title><b>Aktifitas Terbaru</b></v-card-title>
        </div>
        <v-divider class="mx-2"></v-divider>
        <div class="">
          <v-table class="no-divider" style="background-color: #FAFAFA;">
            <thead>
              <tr class="">
                <th style="min-width: 120px;" class="text-left font-weight-bold">
                  Tanggal
                </th>
                <th style="min-width: 100px;" class="text-left font-weight-bold">
                  Pengguna
                </th>
                <th style="min-width: 180px;" class="text-left font-weight-bold">
                  Aktivitas
                </th>
                <th style="min-width: 180px;" class="text-left font-weight-bold">
                  Status
                </th>
              </tr>
            </thead>
              <tbody>
                <tr
                v-for="item in log.slice(0,5)"
                :key="item.name"
                >
              <td>{{ item.created_at }}</td>
              <td>{{ item.username }}</td>
              <td>{{ item.aktivitas }}</td>
              <td>
                <v-badge v-if="item.status=='Menunggu' || item.status=='Belum Ditanggapi'" dot inline color="#FFE642"></v-badge>
                <v-badge v-else-if="item.status=='Draft'" dot inline color="#4BB0EB"></v-badge>
                <v-badge v-else-if="item.status=='Publish' || item.status=='Sudah Ditanggapi'" dot inline color="#4A975B"></v-badge>
                <v-badge v-else-if="item.status=='Ditolak'" dot inline color="#FF4242"></v-badge>{{ item.status }}
              </td>
              </tr>
              </tbody>
          </v-table>
        </div>
      </v-card>
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
      counterAduan:[],
      series: [],
      options: {},
      today : this.todayDate()
    }
  },
  mounted(){
    this.historyLog();
    this.aduanCount();
    this.getAduanStatsTahun()
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
      axios.get('/api/index-log')
      .then(res=>{
        console.log(res.data.data)
        this.log = res.data.data
      })
    },
    aduanCount(){
      axios.get("/api/count-aduan")
      .then(res=>{
        console.log(res.data)
        this.counterAduan = res.data
        if (this.counterAduan.belum_ditanggapi < 9) {
          this.counterAduan.belum_ditanggapi = "0"+res.data.belum_ditanggapi
        }
        if (this.counterAduan.sudah_ditanggapi < 9) {
          this.counterAduan.sudah_ditanggapi = "0"+res.data.sudah_ditanggapi
        }
        if (this.counterAduan.total < 9) {
          this.counterAduan.total = "0"+res.data.total
        }
      })
    },
    getAduanStatsBulan() {
      axios.get("/api/count-aduan-perbulan")
        .then(res => {
          this.series = res.data.series
          this.options = {
            ...this.options,
            ...res.data.options
          }
        })
        .catch(err => {
          console.error("Error fetching stats:", err)
        })
    },
    getAduanStatsTahun() {
      axios.get("/api/count-aduan-pertahun")
        .then(res => {
          this.series = res.data.series
          this.options = {
            ...this.options,
            ...res.data.options
          }
        })
        .catch(err => {
          console.error("Error fetching stats:", err)
        })
    }
  },

}
</script>

<style scooped>

</style>

