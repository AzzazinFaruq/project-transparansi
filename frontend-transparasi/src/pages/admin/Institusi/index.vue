<template>
  <div class="">
  <div class="mt-3">
    <div class="d-flex align-center justify-space-between mr-3 mt-2">
      <v-card-title><b>Daftar Institusi</b></v-card-title>
      <v-btn color="#BF3232" prepend-icon="mdi-plus" href="/admin/institusi/tambah">Institusi</v-btn>
    </div>
  </div>
  <div class="">
          <v-table class="no-divider ">
            <thead style="">
              <tr class="">
                <th class=" font-weight-bold">
                  Tanggal
                </th>
                <th class=" font-weight-bold">
                  Nama Institusi
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
              <td>{{ item.CreatedAt}}</td>
              <td>{{ item.nama_institusi }}</td>
              <td>
                <v-btn class="rounded-lg" style="background-color:#3884B0;color: white;text-transform: none;" @click="detail(item.id)">
                  Lihat Detail
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
  data(){
    return{
      headers:[],
      institusi:[],
      currentPage:1,
      itemsPerPage:10,
    }
  },
  mounted(){
    this.getInstitusi();
  },
  methods:{
    changePage(page) {
      this.currentPage = page;
    },
    getInstitusi(){
      axios.get('/api/index-institusi')
      .then((res)=>{
        this.institusi = res.data.data;
      })
      .catch((err)=>{
        console.log(err);
      })
    }
  },
  computed: {
    totalPages() {
      return Math.ceil(this.institusi.length / this.itemsPerPage);
    },
    paginatedItems() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.institusi.slice(start, end);
    },
  },
}
</script>