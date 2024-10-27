<template>
  <div class="mt-3">
        <div class="d-flex align-center justify-space-between mr-3 mt-2">
          <v-card-title><b>Daftar Program</b></v-card-title>
          <a href="">
            <v-icon class="">mdi-dots-vertical</v-icon>
          </a>
        </div>
        <v-divider class="mx-2"></v-divider>
        <div class="">
          <v-table class="no-divider ">
            <thead style="">
              <tr class="">
                <th class=" font-weight-bold">
                  Tanggal
                </th>
                <th class=" font-weight-bold">
                  Nama
                </th>
                <th class=" font-weight-bold">
                  Status
                </th>
                <th class=" font-weight-bold">
                  Action
                </th>
              </tr>
            </thead>
              <tbody>
                <tr
                v-for="(item, index) in paginatedItems" :key="index"
                >
              <td>{{ item.created_at}}</td>
              <td>{{ item.nama_program }}</td>
              <td>{{ item.status }}</td>
              <td>
                <v-btn class="rounded-lg" style="background-color:#3884B0;color: white;text-transform: none;" @click="detail(item.id)">
                  Detail
                </v-btn>
                <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #387144;color: white;" @click="disetujui(item.id)" icon>
                  <v-icon>mdi-check-bold</v-icon>
                </v-btn>
                <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #BF3232;color: white;" @click="ditolak(item.id)" icon>
                  <b style="font-weight: 900;">X</b>
                </v-btn>
              </td>

              </tr>
              </tbody>
          </v-table>
        </div>
        <v-divider class="my-3"></v-divider>
        <div class="d-flex justify-end align-center">
          <p class="mr-3">Data/Halaman :</p>
          <div class="mr-5" style="width: 100px;height: 40px;">
            <v-select
            density="compact"
            class="custom-outline"
            v-model="itemsPerPage"
            :items="[5, 10, 25, 50, 100]"
            variant="outlined"
          ></v-select>
          </div>
          <p class="mr-5">{{ currentPage }} dari {{ totalPages }} Halaman</p>
          <v-pagination
            class="custom-pagination"
            :total-visible="0"
            border
            v-model="currentPage"
            :length="totalPages"
            @input="changePage"
          >
          </v-pagination>
        </div>
    </div>
</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      detailProgram:[],
      Userlist:[],
      currentPage: 1,
      itemsPerPage: 10,
    }
  },
  mounted() {
    this.user();
  },
  methods: {
    user(){
      axios.all([
        axios.get("/api/program/disetujui"),
        axios.get("/api/program/ditolak")
      ])
      .then(axios.spread((resDisetujui, resDitolak) => {
        // Mengambil data dari kedua respons
        const dataDisetujui = resDisetujui.data.data; // Asumsi struktur data
        const dataDitolak = resDitolak.data.data; // Asumsi struktur data

        // Menggabungkan kedua array ke dalam Userlist
        this.Userlist = [...dataDisetujui, ...dataDitolak];

        console.log(this.Userlist); // Tampilkan hasil gabungan
      }))
      .catch(error => {
        console.error('Error fetching data:', error);
      });
    },
    changePage(page) {
      this.currentPage = page;
    },
    detail(id){

      axios.get(`/api/program/${id}`)
      .then(res=>{
        this.detailProgram = res.data.data
      })
    },
    disetujui(id){
      console.log(id)
      axios.get(`/api/program/accept/${id}`)
      .then(res=>{
        console.log(res.data.data)
        this.user();
      })
    },
    ditolak(id){
      console.log(id)
      axios.get(`/api/program/reject/${id}`)
      .then(res=>{
        console.log(res.data.data)
        this.user();
      })
    }
  },
  computed: {
    totalPages() {
      return Math.ceil(this.Userlist.length / this.itemsPerPage);
    },
    paginatedItems() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.Userlist.slice(start, end);
    },
  },
}
</script>
