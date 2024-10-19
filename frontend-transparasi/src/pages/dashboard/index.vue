<template>
  <v-container >
    <div class="d-flex justify-space-between">
      <div class="header-page">
        <h1 class="title">Dashboard</h1>
      <p class="caption">Home > Dashboard</p>
      </div>
      <div class="">
        {{ today }}
      </div>
    </div>

    <div class="mt-5">
      <v-row>
        <v-col cols="12" sm="6" md="4" lg="3">
          <card-dashboard
          name="Total Dana"
          icon="mdi-cash"
          classes="icon-container background-green"
          />
        </v-col>
        <v-col cols="12" sm="6" md="4" lg="3">
          <card-dashboard
          name="Dana Tersedia"
          icon="mdi-cash"
          classes="icon-container background-blue"
          />
        </v-col>
        <v-col cols="12" sm="6" md="4" lg="3">
          <card-dashboard
          name="Dana Terpakai"
          icon="mdi-cash"
          classes="icon-container background-red"
          />
        </v-col>
        <v-col cols="12" sm="6" md="4" lg="3">
          <card-dashboard
          name="Dana Terkunci"
          icon="mdi-cash"
          classes="icon-container background-yellow"
          />
        </v-col>
      </v-row>
    </div>

    <div class="mt-3">
      <v-card>
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
                  Program
                </th>
                <th class="text-left font-weight-bold">
                  Status
                </th>
              </tr>
            </thead>
          </v-table>
        </div>
      </v-card>
    </div>

  </v-container>
</template>
<script>
import cardDashboard from '@/components/card-dashboard.vue';

export default{
  components: { cardDashboard },
  data() {
    return {
      today : this.todayDate()
    }
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
    }
  },
}
</script>

<style scooped>

</style>
