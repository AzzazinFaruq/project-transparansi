<template>
  <div class="mt-3">
    <v-card class="mx-4">
        <div class="d-flex align-center justify-space-between mr-3 mt-2">
          <v-card-title><b>Daftar Keluhan</b></v-card-title>
          <a href="">
            <v-icon class="">mdi-dots-vertical</v-icon>
          </a>
        </div>
        <v-divider class="mx-2"></v-divider>
        <div class="">
          <v-table class="no-divider ">
            <thead style="">
              <tr class="">
                <th style="min-width: 100px;" class=" font-weight-bold">
                  Tanggal
                </th>
                <th style="min-width: 150px;" class=" font-weight-bold">
                  Nama
                </th>
                <th style="min-width: 150px;" class=" font-weight-bold">
                  Status
                </th>
                <th style="min-width: 100px;" class=" font-weight-bold">
                  Aksi
                </th>
              </tr>
            </thead>
              <tbody>
                <tr
                v-for="(item, index) in paginatedItems" :key="index"
                >
              <td>{{ item.created_at}}</td>
              <td>{{ item.program.nama_program }}</td>
              <td>
                <v-badge v-if="item.status=='Belum Ditanggapi'" dot inline color="#FFE642"></v-badge>
                <v-badge v-else-if="item.status=='Sudah Ditanggapi'" dot inline color="#4A975B"></v-badge>
                {{ item.status }}
              </td>
              <td>
                <v-btn class="rounded-lg" style="background-color:#3884B0;color: white;text-transform: none;" @click="detailAduan(item.id)">
                  Detail
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
      </v-card>
    </div>
</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      headers: [
        { title: "Nama", value: "username" },
        { title: "Jabatan", value: "" },
        { title: "Email", value: "email" },
        { title: "No Hp", value: "no_hp" },
        { title: "Action", value: "actions" },
      ],
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
      axios.get("/api/index-aduan")
      .then(res=>{
        console.log(res.data.data)
        this.Userlist=res.data.data
      })
    },
    changePage(page) {
      this.currentPage = page;
    },
    detailAduan(id){
      this.$router.push(`/dprd/keluhan/${id}`)
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
