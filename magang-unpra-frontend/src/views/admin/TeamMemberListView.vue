<template>
  <div class="min-h-screen bg-gray-100 flex">
    <!-- Sidebar -->
    <aside class="w-64 bg-gray-950 text-white flex flex-col fixed h-full z-40">
      <div class="p-6 border-b border-gray-800">
        <div class="flex items-center gap-3">
          <div class="w-9 h-9 bg-green-600 rounded-xl flex items-center justify-center font-black text-sm">T</div>
          <div>
            <p class="font-black text-sm">TELPP</p>
            <p class="text-xs text-gray-500">Management Profile</p>
          </div>
        </div>
      </div>
      <nav class="flex-1 p-4 space-y-1">
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-2 px-2">OVERVIEW</p>
        <RouterLink to="/admin/dashboard"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📊 Dashboard
        </RouterLink>
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-4 px-2">CONTENT ENGINE</p>
        <RouterLink to="/admin/news"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📰 Corporate News
        </RouterLink>
        <RouterLink to="/admin/slider"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🖼️ Product Slider
        </RouterLink>
        <RouterLink to="/admin/product-page"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📝 Product Page
        </RouterLink>
        <RouterLink to="/admin/about"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🏢 About Section
        </RouterLink>
        <RouterLink to="/admin/our-company"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🏛️ Our Company
        </RouterLink>
        <RouterLink to="/admin/team-members"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
          👥 Our Team
        </RouterLink>
        <RouterLink to="/admin/menus"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📋 Menu
        </RouterLink>
      </nav>
      <div class="p-4 border-t border-gray-800">
        <button @click="logout" class="text-sm text-gray-400 hover:text-white transition px-2">→ Logout</button>
      </div>
    </aside>

    <!-- Main -->
    <main class="flex-1 ml-64 p-10">
      <!-- Header -->
      <div class="flex justify-between items-center mb-6">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">Our Team</h1>
          <p class="text-gray-400 text-sm mt-1">Kelola Dewan Direksi dan Struktur Organisasi pada halaman /our-team.</p>
        </div>
        <button v-if="activeTab === 'direksi'" @click="openCreate"
          class="bg-green-600 text-white px-6 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition flex items-center gap-2">
          <span class="text-lg leading-none">+</span> Tambah Anggota
        </button>
        <div v-else class="flex items-center gap-2">
          <p class="text-sm text-gray-400">Template struktur sudah tersedia — tinggal ganti foto, nama, dan jabatan.</p>
        </div>
      </div>

      <!-- Tabs -->
      <div class="flex items-center gap-1 mb-6 bg-gray-200/70 p-1 rounded-xl w-fit">
        <button
          @click="activeTab = 'direksi'"
          :class="['px-5 py-2 rounded-lg text-sm font-semibold transition',
            activeTab === 'direksi' ? 'bg-white text-gray-900 shadow-sm' : 'text-gray-500 hover:text-gray-700']">
          Dewan Direksi
        </button>
        <button
          @click="activeTab = 'struktur'"
          :class="['px-5 py-2 rounded-lg text-sm font-semibold transition',
            activeTab === 'struktur' ? 'bg-white text-gray-900 shadow-sm' : 'text-gray-500 hover:text-gray-700']">
          Struktur Organisasi
        </button>
      </div>

      <!-- Alert -->
      <div v-if="alert.show"
        :class="['mb-4 px-4 py-3 rounded-xl text-sm font-medium flex items-center gap-2',
          alert.type === 'success' ? 'bg-green-50 text-green-700 border border-green-200' : 'bg-red-50 text-red-700 border border-red-200']">
        {{ alert.type === 'success' ? '✅' : '❌' }} {{ alert.message }}
      </div>

      <!-- ============ TAB: DEWAN DIREKSI ============ -->
      <div v-if="activeTab === 'direksi'" class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <div v-if="loading" class="flex items-center justify-center py-16">
          <div class="w-5 h-5 border-2 border-gray-200 border-t-green-500 rounded-full animate-spin"></div>
        </div>

        <table v-else class="w-full text-sm">
          <thead class="bg-gray-50 text-gray-500 text-xs font-semibold tracking-wider">
            <tr>
              <th class="text-left px-6 py-4">Foto</th>
              <th class="text-left px-6 py-4">Nama</th>
              <th class="text-left px-6 py-4">Jabatan</th>
              <th class="text-left px-6 py-4">Petinggi</th>
              <th class="text-left px-6 py-4">Urutan</th>
              <th class="text-left px-6 py-4">Status</th>
              <th class="text-right px-6 py-4">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="m in members" :key="m.ID" class="border-t border-gray-100 hover:bg-gray-50 transition">
              <td class="px-6 py-4">
                <div class="w-12 h-12 rounded-full overflow-hidden bg-gray-100 border border-gray-200">
                  <img v-if="m.photo_path" :src="getImageUrl(m.photo_path)" class="w-full h-full object-cover" />
                  <div v-else class="w-full h-full flex items-center justify-center text-gray-400 text-lg">👤</div>
                </div>
              </td>
              <td class="px-6 py-4 font-semibold text-gray-900">{{ m.name }}</td>
              <td class="px-6 py-4 text-gray-600">{{ m.position }}</td>
              <td class="px-6 py-4">
                <span v-if="m.is_featured" class="px-2 py-1 rounded-lg text-xs font-semibold bg-yellow-100 text-yellow-700">Petinggi</span>
              </td>
              <td class="px-6 py-4 text-gray-500">{{ m.sort_order }}</td>
              <td class="px-6 py-4">
                <span :class="['px-2.5 py-1 rounded-lg text-xs font-semibold',
                  m.is_active ? 'bg-green-100 text-green-700' : 'bg-gray-100 text-gray-500']">
                  {{ m.is_active ? 'Aktif' : 'Nonaktif' }}
                </span>
              </td>
              <td class="px-6 py-4 text-right">
                <div class="flex items-center justify-end gap-2">
                  <button @click="openEdit(m)"
                    class="px-3 py-1.5 rounded-lg text-xs font-semibold bg-blue-50 text-blue-600 hover:bg-blue-100 transition">
                    Edit
                  </button>
                  <button @click="remove(m.ID)"
                    class="px-3 py-1.5 rounded-lg text-xs font-semibold bg-red-50 text-red-600 hover:bg-red-100 transition">
                    Hapus
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="!members.length">
              <td colspan="7" class="px-6 py-16 text-center text-gray-400">
                <p class="text-4xl mb-3">👥</p>
                <p>Belum ada anggota tim.</p>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- ============ TAB: STRUKTUR ORGANISASI (gambar bagan) ============ -->
      <div v-else>
        <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-8">
          <h3 class="text-lg font-bold text-gray-900 mb-2">Upload Bagan Struktur Organisasi</h3>
          <p class="text-sm text-gray-500 mb-6">Admin membuat bagan struktur di aplikasi lain (Canva, Word, dll), screenshot, lalu upload di sini. Satu gambar saja.</p>

          <div class="flex items-center gap-6">
            <div v-if="orgChartImage" class="w-48 h-48 rounded-xl overflow-hidden border border-gray-200 bg-gray-50 flex-shrink-0">
              <img :src="getImageUrl(orgChartImage)" class="w-full h-full object-cover" />
            </div>
            <div class="flex-1">
              <input type="file" accept="image/*" ref="chartInput" class="hidden" @change="handleChartUpload" />
              <button @click="$refs.chartInput.click()"
                class="border border-dashed border-gray-300 rounded-xl px-5 py-4 text-sm text-gray-500 hover:border-green-500 hover:text-green-600 transition w-full text-left">
                {{ uploadingChart ? 'Mengupload...' : 'Pilih gambar bagan dari komputer' }}
              </button>
              <p class="text-xs text-gray-400 mt-2">JPG, PNG, WebP. Maks 5MB.</p>
              <div v-if="orgChartImage" class="flex gap-2 mt-3">
                <button @click="saveChart" :disabled="savingChart"
                  class="bg-green-600 text-white px-4 py-2 rounded-xl text-sm font-semibold hover:bg-green-700 transition disabled:opacity-50">
                  {{ savingChart ? 'Menyimpan...' : 'Simpan ke Website' }}
                </button>
                <button @click="removeChart"
                  class="border border-red-200 text-red-600 px-4 py-2 rounded-xl text-sm font-semibold hover:bg-red-50 transition">
                  Hapus Gambar
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- ============ MODAL: CREATE/EDIT ANGGOTA DIREKSI ============ -->
      <Transition name="modal-fade">
        <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 backdrop-blur-sm p-4"
          @click.self="showModal = false">
          <div class="bg-white rounded-2xl w-full max-w-lg p-8 shadow-2xl">
            <h2 class="text-xl font-black text-gray-900 mb-6">{{ editing ? 'Edit Anggota' : 'Tambah Anggota' }}</h2>

            <div class="space-y-4">
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-1.5">Nama Lengkap</label>
                <input v-model="form.name" type="text" placeholder="Contoh: Budi Santoso"
                  class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
              </div>
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-1.5">Jabatan</label>
                <input v-model="form.position" type="text" placeholder="Contoh: Direktur Utama"
                  class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
              </div>
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-1.5">Deskripsi / Bio</label>
                <textarea v-model="form.description" rows="4" placeholder="Deskripsi singkat tentang anggota tim..."
                  class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-500 resize-none"></textarea>
              </div>
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-1.5">Foto</label>
                <div class="flex items-center gap-4">
                  <div v-if="form.photo_path" class="w-20 h-20 rounded-xl overflow-hidden bg-gray-100 border border-gray-200 flex-shrink-0">
                    <img :src="getImageUrl(form.photo_path)" class="w-full h-full object-cover" />
                  </div>
                  <div class="flex-1">
                    <input type="file" accept="image/*" ref="photoInput" class="hidden" @change="handlePhoto" />
                    <button @click="$refs.photoInput.click()"
                      class="border border-dashed border-gray-300 rounded-xl px-4 py-3 text-sm text-gray-500 hover:border-green-500 hover:text-green-600 transition w-full text-left">
                      {{ uploadingPhoto ? 'Mengupload...' : 'Pilih foto dari komputer' }}
                    </button>
                    <p class="text-xs text-gray-400 mt-1">JPG, PNG, WebP. Maks 5MB.</p>
                  </div>
                </div>
              </div>
              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="text-sm font-semibold text-gray-700 block mb-1.5">Urutan Tampil</label>
                  <input v-model.number="form.sort_order" type="number" min="0"
                    class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
                </div>
                <div class="space-y-3">
                  <div>
                    <label class="text-sm font-semibold text-gray-700 block mb-1.5">Status</label>
                    <label class="flex items-center gap-3 mt-2 cursor-pointer">
                      <input type="checkbox" v-model="form.is_active"
                        class="w-5 h-5 rounded-lg border-gray-300 text-green-600 focus:ring-green-500" />
                      <span class="text-sm text-gray-700">Aktif</span>
                    </label>
                  </div>
                  <div>
                    <label class="flex items-center gap-3 cursor-pointer">
                      <input type="checkbox" v-model="form.is_featured"
                        class="w-5 h-5 rounded-lg border-gray-300 text-yellow-500 focus:ring-yellow-500" />
                      <span class="text-sm text-gray-700 font-semibold">Petinggi / Direktur Utama</span>
                    </label>
                    <p class="text-xs text-gray-400 mt-1">Anggota yang dicentang akan tampil dengan kartu besar di halaman publik.</p>
                  </div>
                </div>
              </div>
            </div>

            <div class="flex gap-3 mt-6 pt-4 border-t border-gray-100">
              <button @click="save" :disabled="saving"
                class="bg-green-600 text-white px-6 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition disabled:opacity-50 flex items-center gap-2">
                <div v-if="saving" class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                {{ saving ? 'Menyimpan...' : 'Simpan' }}
              </button>
              <button @click="showModal = false"
                class="border border-gray-200 px-6 py-2.5 rounded-xl text-sm font-semibold text-gray-600 hover:bg-gray-50 transition">
                Batal
              </button>
            </div>
          </div>
        </div>
      </Transition>


    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import api from '../../services/api'

const BASE_URL = 'http://localhost:8080'
const router = useRouter()

/* ── shared ── */
const activeTab = ref('direksi') // 'direksi' | 'struktur'
const alert = ref({ show: false, type: 'success', message: '' })

const getImageUrl = (path) => {
  if (!path) { return '' }
  if (path.startsWith('http')) { return path }
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const showAlert = (type, message) => {
  alert.value = { show: true, type, message }
  setTimeout(() => { alert.value.show = false }, 3000)
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}

/* =========================================================
   TAB 1 — DEWAN DIREKSI (sama seperti sebelumnya)
   ========================================================= */
const loading = ref(true)
const saving = ref(false)
const uploadingPhoto = ref(false)
const showModal = ref(false)
const editing = ref(false)
const editingId = ref(null)
const members = ref([])
const photoInput = ref(null)

const form = ref({
  name: '',
  position: '',
  description: '',
  photo_path: '',
  sort_order: 0,
  is_active: true,
  is_featured: false
})

const fetchMembers = async () => {
  loading.value = true
  try {
    const res = await api.get('/admin/team-members')
    members.value = res.data?.data || []
  } catch { showAlert('error', 'Gagal memuat data') }
  finally { loading.value = false }
}

const openCreate = () => {
  editing.value = false
  editingId.value = null
  form.value = { name: '', position: '', description: '', photo_path: '', sort_order: 0, is_active: true, is_featured: false }
  showModal.value = true
}

const openEdit = (m) => {
  editing.value = true
  editingId.value = m.ID
  form.value = {
    name: m.name,
    position: m.position,
    description: m.description || '',
    photo_path: m.photo_path || '',
    sort_order: m.sort_order || 0,
    is_active: m.is_active,
    is_featured: m.is_featured
  }
  showModal.value = true
}

const handlePhoto = async (e) => {
  const file = e.target.files?.[0]
  if (!file) {return}
  if (file.size > 5 * 1024 * 1024) {
    showAlert('error', 'Ukuran file maksimal 5MB')
    return
  }
  uploadingPhoto.value = true
  try {
    const fd = new FormData()
    fd.append('image', file)
    const res = await api.post('/upload', fd, { headers: { 'Content-Type': 'multipart/form-data' } })
    form.value.photo_path = `uploads/${res.data?.filename || ''}`
    showAlert('success', 'Foto berhasil diupload!')
  } catch { showAlert('error', 'Gagal upload foto') }
  finally {
    uploadingPhoto.value = false
    if (photoInput.value) {photoInput.value.value = ''}
  }
}

const save = async () => {
  if (!form.value.name) { showAlert('error', 'Nama wajib diisi!'); return }
  if (!form.value.position) { showAlert('error', 'Jabatan wajib diisi!'); return }
  saving.value = true
  try {
    if (editing.value) {
      await api.put(`/admin/team-members/${editingId.value}`, form.value)
      showAlert('success', 'Data berhasil diupdate!')
    } else {
      await api.post('/admin/team-members', form.value)
      showAlert('success', 'Anggota berhasil ditambahkan!')
    }
    showModal.value = false
    await fetchMembers()
  } catch (e) {
    showAlert('error', e.response?.data?.error || 'Gagal menyimpan')
  } finally { saving.value = false }
}

const remove = async (id) => {
  if (!confirm('Yakin ingin menghapus anggota ini?')) {return}
  try {
    await api.delete(`/admin/team-members/${id}`)
    showAlert('success', 'Berhasil dihapus!')
    await fetchMembers()
  } catch { showAlert('error', 'Gagal menghapus') }
}

/* =========================================================
   TAB 2 — STRUKTUR ORGANISASI (gambar bagan)
   ========================================================= */
const orgChartImage = ref('')
const uploadingChart = ref(false)
const savingChart = ref(false)
const chartInput = ref(null)

const fetchOrgChart = async () => {
  try {
    const res = await api.get('/org-chart')
    orgChartImage.value = res.data?.data?.image_path || ''
  } catch (_e) { /* org-chart optional */ }
}

const handleChartUpload = async (e) => {
  const file = e.target.files?.[0]
  if (!file) {return}
  if (file.size > 5 * 1024 * 1024) {
    showAlert('error', 'Ukuran file maksimal 5MB')
    return
  }
  uploadingChart.value = true
  try {
    const fd = new FormData()
    fd.append('image', file)
    const res = await api.post('/upload', fd, { headers: { 'Content-Type': 'multipart/form-data' } })
    orgChartImage.value = `uploads/${res.data?.filename || ''}`
    showAlert('success', 'Gambar berhasil diupload! Silakan klik Simpan ke Website.')
  } catch (e) {
    showAlert('error', e.response?.data?.error || e.message || 'Gagal upload gambar')
  }
  finally {
    uploadingChart.value = false
    if (chartInput.value) {chartInput.value.value = ''}
  }
}

const saveChart = async () => {
  if (!orgChartImage.value) { showAlert('error', 'Tidak ada gambar untuk disimpan'); return }
  savingChart.value = true
  try {
    await api.put('/admin/org-chart', { image_path: orgChartImage.value })
    showAlert('success', 'Bagan struktur organisasi berhasil disimpan!')
  } catch (e) {
    showAlert('error', e.response?.data?.error || e.message || 'Gagal menyimpan')
  }
  finally { savingChart.value = false }
}

const removeChart = async () => {
  if (!confirm('Hapus gambar bagan struktur organisasi?')) {return}
  try {
    await api.put('/admin/org-chart', { image_path: '' })
    orgChartImage.value = ''
    showAlert('success', 'Gambar berhasil dihapus!')
  } catch { showAlert('error', 'Gagal menghapus gambar') }
}

/* =========================================================
   INIT
   ========================================================= */
onMounted(() => {
  if (!localStorage.getItem('token')) { router.push('/admin/login'); return }
  fetchMembers()
  fetchOrgChart()
})
</script>

<style scoped>
.modal-fade-enter-active, .modal-fade-leave-active { transition: opacity 0.25s ease; }
.modal-fade-enter-from, .modal-fade-leave-to { opacity: 0; }
</style>