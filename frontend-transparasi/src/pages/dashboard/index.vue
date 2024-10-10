<template>
  <v-container>
    <div class="d-flex justify-space-between">
      <div class="header-page">
        <h1 class="title">Dashboard</h1>
      <p class="caption">Home > Dashboard</p>
      </div>
      <div class="">
        {{ today }}
      </div>
    </div>
  </v-container>
</template>
<script>
import router from '@/router';
import { authStore } from '@/store/auth';
const auth = authStore();
auth.check(router);

export default{
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
.header-page{
  .title{
    font-size: 24px;
  }
  .caption{
    font-size: 16px;
  }
}
</style>
