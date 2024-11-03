<template>
  <v-container>
  <div class="mt-3">
        <div class="d-flex align-center justify-space-between mr-3 mt-2">
          <v-card-title><b>Daftar Keluhan</b></v-card-title>
          <div class="d-flex align-center">
           <p class="mr-3">Filter : </p>
           <div style=" min-width: 120px;">
            <v-btn
                variant="text"
                style="border-color: #BF3232;  text-transform: none; letter-spacing: 0.5px; margin-top: 3px; padding: 0;"
                append-icon="mdi-chevron-down"
              >
                {{ filter }}
                <v-menu activator="parent">
                  <v-list>
                    <v-list-item
                    v-for="item in filterList"
                    :key="item"
                    @click="selectedFilter(item);"
                    >
                      <v-list-item-title >{{ item }}</v-list-item-title>
                    </v-list-item>
                  </v-list>
                </v-menu>
              </v-btn>
           </div>
          </div>
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
              <td
              style="min-width: 150px;"
              >
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
    </div>
  </v-container>
</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      filter:'All Status',
      filterList:["All Status","Belum Ditanggapi","Sudah Ditanggapi"],
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
    selectedFilter(item){
      this.filter = item
      if (this.filter == "All Status") {
        axios.get("/api/index-aduan")
        .then(res=>{
          this.Userlist = res.data.data
        })
      } else {
        axios.get(`/api/aduan/status/${item}`)
        .then(res=>{
          this.Userlist = res.data.data
        })
      }
    },

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
      this.$router.push(`/admin/keluhan/${id}`)
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
