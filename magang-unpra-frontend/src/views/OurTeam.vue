<template>
  <div class="page">

    <!-- HERO -->
    <PageHero
      title="Our Team"
      subtitle="Meet the talented people who helped make PT TELPP what it is today."
      :breadcrumbs="[{ label: 'Home', to: '/' }, { label: 'Our Company', to: '/our-company' }, { label: 'Our Team' }]"
    />

    <!-- TEAM LIST -->
    <section class="team-section">
      <div class="section-inner">
        <h2 class="section-title">Dewan Direksi</h2>

        <div v-if="loading" class="flex items-center justify-center py-24">
          <div class="w-6 h-6 border-2 border-emerald-600 border-t-transparent rounded-full animate-spin"></div>
        </div>

        <div v-if="featuredMembers.length" class="featured-list">
          <div v-for="featured in featuredMembers" :key="featured.ID" class="featured-card">
            <div class="featured-photo">
              <span class="featured-photo-placeholder">👤</span>
              <img v-if="featured.photo_path" :src="getImageUrl(featured.photo_path)" :alt="featured.name" />
            </div>
            <div class="featured-info">
              <h3 class="featured-name">{{ featured.name }}</h3>
              <p class="featured-pos">{{ featured.position }}</p>
              <p class="featured-desc">{{ featured.description }}</p>
              <button @click="openDetail(featured)" class="featured-btn">Selengkapnya →</button>
            </div>
          </div>
        </div>

        <div v-if="regulars.length" class="team-grid">
          <div
            v-for="(m, i) in regulars"
            :key="m.ID"
            class="team-card"
            :style="{ animationDelay: (i * 0.12) + 's' }"
            @click="openDetail(m)"
          >
            <div class="team-photo">
              <div class="photo-placeholder">👤</div>
              <img v-if="m.photo_path" :src="getImageUrl(m.photo_path)" :alt="m.name" class="team-img" />
            </div>
            <div class="team-info">
              <h3 class="team-name">{{ m.name }}</h3>
              <p class="team-position">{{ m.position }}</p>
            </div>
          </div>
        </div>

        <div v-else class="empty-state">
          <p class="empty-icon">👥</p>
          <p class="empty-text">Belum ada data anggota tim.</p>
        </div>
      </div>
    </section>

    <!-- MODAL DETAIL ANGGOTA -->
    <Transition name="modal">
      <div v-if="showDetail" class="modal-overlay" @click.self="showDetail = false">
        <!-- MODAL: Featured / Petinggi (horizontal besar) -->
        <div v-if="selectedMember?.is_featured" class="modal-card modal-featured">
          <button @click="showDetail = false" class="modal-x">×</button>
          <div class="mf-photo">
            <span v-if="!selectedMember?.photo_path" class="mf-photo-placeholder">👤</span>
            <img v-if="selectedMember?.photo_path" :src="getImageUrl(selectedMember.photo_path)" :alt="selectedMember.name" />
            <div class="mf-badge">PETINGGI</div>
          </div>
          <div class="mf-info">
            <h3 class="mf-name">{{ selectedMember?.name }}</h3>
            <p class="mf-pos">{{ selectedMember?.position }}</p>
            <div class="mf-divider"></div>
            <p class="mf-desc">{{ selectedMember?.description || 'Tidak ada deskripsi.' }}</p>
          </div>
        </div>

        <!-- MODAL: Regular -->
        <div v-else class="modal-card modal-regular">
          <button @click="showDetail = false" class="modal-x">×</button>
          <div class="mr-photo">
            <span v-if="!selectedMember?.photo_path" class="modal-photo-placeholder">👤</span>
            <img v-if="selectedMember?.photo_path" :src="getImageUrl(selectedMember.photo_path)" :alt="selectedMember.name" />
          </div>
          <div class="mr-info">
            <h3 class="mr-name">{{ selectedMember?.name }}</h3>
            <p class="mr-pos">{{ selectedMember?.position }}</p>
            <p class="mr-desc">{{ selectedMember?.description || 'Tidak ada deskripsi.' }}</p>
          </div>
        </div>
      </div>
    </Transition>

    <!-- STRUKTUR ORGANISASI (gambar bagan) -->
    <section class="orgchart-section">
      <div class="section-inner">
        <h2 class="section-title">Struktur Organisasi</h2>
        <div v-if="orgChartImage" class="orgchart-img-wrap">
          <img :src="getImageUrl(orgChartImage)" alt="Struktur Organisasi" class="orgchart-img" />
        </div>
        <div v-else class="empty-state">
          <p class="empty-icon">🏢</p>
          <p class="empty-text">Struktur organisasi belum tersedia.</p>
        </div>
      </div>
    </section>

    <footer class="site-footer">
      <div class="footer-container">
        <div class="footer-left-content"></div>
        <div class="footer-copyright">
          <p>Copyright 2026 PT TELPP. All right reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import api from '../services/api'
import PageHero from '../components/PageHero.vue'

const BASE_URL = import.meta.env.VITE_BASE_URL || 'http://localhost:8080'
const members = ref([])
const loading = ref(true)
const showDetail = ref(false)
const selectedMember = ref(null)

const featuredMembers = computed(() => members.value.filter(m => m.is_featured))
const regulars = computed(() => members.value.filter(m => !m.is_featured))

const getImageUrl = (path) => {
  if (!path) {return ''}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const openDetail = (m) => {
  selectedMember.value = m
  showDetail.value = true
}

const orgChartImage = ref('')

onMounted(async () => {
  try {
    const res = await api.get('/team-members')
    members.value = res.data?.data || []
  } catch {
    members.value = []
  } finally {
    loading.value = false
  }
  try {
    const res = await api.get('/org-chart')
    orgChartImage.value = res.data?.data?.image_path || ''
  } catch (_e) { /* org-chart optional */ }
})
</script>

<style scoped>
* { box-sizing: border-box; margin: 0; padding: 0; }
.page {
  font-family: 'Segoe UI', system-ui, Arial, sans-serif;
  color: #1a1a1a;
  background: #fff;
  overflow-x: hidden;
}

/* ── HERO ── */
.hero {
  position: relative;
  min-height: 400px;
  display: flex;
  align-items: flex-end;
  overflow: hidden;
}
.hero-bg { position: absolute; inset: 0; z-index: 0; }
.hero-img {
  width: 100%; height: 100%; object-fit: cover;
  object-position: center 40%;
  display: block;
  transform: scale(1.06);
  animation: heroZoom 14s ease-out forwards;
}
@keyframes heroZoom { to { transform: scale(1); } }

.hero-overlay {
  position: absolute; inset: 0;
  background: linear-gradient(105deg, rgba(5,28,14,.88) 0%, rgba(5,28,14,.6) 42%, rgba(5,28,14,.18) 100%);
}

.hero-content {
  position: relative; z-index: 2;
  padding: 64px 72px 72px;
  max-width: 660px;
}

.eyebrow {
  font-size: .7rem; font-weight: 700; letter-spacing: .16em;
  color: #5ecb7a; margin-bottom: 18px;
  opacity: 0; transform: translateY(14px);
  transition: opacity .6s, transform .6s;
}
.eyebrow.in { opacity: 1; transform: none; }

h1 { margin: 0 0 22px; color: #fff; }
.line {
  display: block;
  font-size: 2.8rem; font-weight: 800; line-height: 1.18;
  opacity: 0; transform: translateY(22px) skewY(1.5deg);
  transition: opacity .7s cubic-bezier(.22,1,.36,1), transform .7s cubic-bezier(.22,1,.36,1);
}
.in .line { opacity: 1; transform: none; }

.accent-line {
  width: 0; height: 3px; background: #2d7a3d;
  border-radius: 2px; margin-bottom: 22px;
  transition: width .8s cubic-bezier(.22,1,.36,1);
}
.accent-line.in { width: 56px; }

.hero-p {
  font-size: .93rem; color: rgba(255,255,255,.8); line-height: 1.8;
  opacity: 0; transform: translateY(14px);
  transition: opacity .7s, transform .7s;
}
.hero-p.in { opacity: 1; transform: none; }

/* ── TEAM SECTION ── */
.team-section {
  padding: 72px 72px 88px;
  background: #fff;
}
.section-inner { max-width: 1200px; margin: 0 auto; }

.section-title {
  font-size: 1.6rem;
  font-weight: 800;
  color: #111827;
  margin-bottom: 32px;
  padding-bottom: 12px;
  border-bottom: 2px solid #e5e7eb;
}

/* ── FEATURED CARD ── */
.featured-list {
  display: flex;
  flex-direction: column;
  gap: 28px;
  margin-bottom: 64px;
}
.featured-card {
  display: flex;
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 16px;
  padding: 0;
  transition: box-shadow .3s, border-color .3s;
  position: relative;
  overflow: hidden;
}
.featured-card:hover {
  border-color: #d1d5db;
  box-shadow: 0 8px 32px rgba(0,0,0,.08);
}
.featured-photo {
  width: 340px;
  flex-shrink: 0;
  position: relative;
  overflow: hidden;
  background: #f3f4f6;
}
.featured-photo img {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: contain;
  object-position: top center;
  background: #f3f4f6;
  transition: transform .5s;
}
.featured-card:hover .featured-photo img { transform: none; }
.featured-photo-placeholder {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 4rem;
  color: #d1d5db;
  background: #f3f4f6;
}
.featured-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 36px 40px;
  min-width: 0;
}
.featured-name {
  font-size: 1.6rem;
  font-weight: 700;
  color: #111827;
  margin: 0 0 6px;
  line-height: 1.2;
}
.featured-pos {
  font-size: .95rem;
  font-weight: 600;
  color: #2563eb;
  margin: 0 0 16px;
}
.featured-desc {
  font-size: .88rem;
  color: #6b7280;
  line-height: 1.75;
  margin: 0 0 20px;
}
.featured-btn {
  align-self: flex-start;
  padding: 10px 28px;
  border-radius: 8px;
  font-size: .85rem;
  font-weight: 600;
  background: #1e3a5f;
  color: #fff;
  border: none;
  cursor: pointer;
  transition: background .2s;
}
.featured-btn:hover { background: #0f2744; }

/* ── GRID LAYOUT ── */
.team-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 28px;
}
.team-card {
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 14px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  transition: box-shadow .3s, border-color .3s, transform .3s;
  animation: fadeInUp .6s cubic-bezier(.22,1,.36,1) both;
  cursor: pointer;
}
.team-card:hover {
  border-color: #d1d5db;
  box-shadow: 0 8px 28px rgba(0,0,0,.1);
  transform: translateY(-4px);
}
@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: none; }
}
.team-photo {
  width: 100%;
  aspect-ratio: 4 / 5;
  overflow: hidden;
  background: #f3f4f6;
  position: relative;
}
.team-photo .team-img {
  position: absolute;
  inset: 0;
  width: 100%; height: 100%; object-fit: contain;
  object-position: top center;
  background: #f3f4f6;
  transition: transform .5s cubic-bezier(.22,1,.36,1);
}
.team-card:hover .team-photo .team-img { transform: none; }
.photo-placeholder {
  position: absolute;
  inset: 0;
  display: flex; align-items: center; justify-content: center;
  font-size: 3rem; color: #d1d5db;
}
.team-info {
  padding: 16px 18px 20px;
  text-align: left;
}
.team-name {
  font-size: 1rem;
  font-weight: 700;
  color: #111827;
  margin-bottom: 3px;
  line-height: 1.3;
}
.team-position {
  font-size: .82rem;
  color: #2563eb;
  font-weight: 500;
  margin-bottom: 8px;
}
.team-desc {
  font-size: .8rem;
  color: #6b7280;
  line-height: 1.55;
  margin-bottom: 8px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.read-more-btn {
  background: none;
  border: none;
  color: #2563eb;
  font-weight: 600;
  font-size: .78rem;
  cursor: pointer;
  padding: 0;
}
.read-more-btn:hover { text-decoration: underline; }

/* ── MODAL ── */
.modal-enter-active, .modal-leave-active { transition: opacity .25s ease; }
.modal-enter-from, .modal-leave-to { opacity: 0; }

.modal-overlay {
  position: fixed;
  inset: 0;
  z-index: 999;
  background: rgba(0,0,0,.55);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}
.modal-card {
  background: #fff;
  border-radius: 14px;
  overflow: hidden;
  position: relative;
  box-shadow: 0 24px 64px rgba(0,0,0,.2);
  max-height: 86vh;
}
.modal-x {
  position: absolute;
  top: 10px; right: 10px;
  width: 28px; height: 28px;
  border-radius: 50%;
  border: none;
  background: #111827;
  color: #fff;
  font-size: 1rem;
  font-weight: 700;
  cursor: pointer;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background .2s;
  line-height: 1;
}
.modal-x:hover { background: #dc2626; }

/* ── FEATURED MODAL ── */
.modal-featured {
  max-width: 800px;
  width: 100%;
  display: flex;
  height: 400px;
}
.mf-photo {
  width: 340px;
  flex-shrink: 0;
  background: #f3f4f6;
  position: relative;
  overflow: hidden;
}
.mf-photo img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  object-position: top center;
  background: #f3f4f6;
}
.mf-photo-placeholder {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 4rem;
  color: #d1d5db;
}
.mf-badge {
  position: absolute;
  top: 14px; left: 14px;
  padding: 4px 12px;
  border-radius: 5px;
  font-size: .68rem;
  font-weight: 700;
  background: #fef3c7;
  color: #92400e;
  text-transform: uppercase;
  letter-spacing: .04em;
}
.mf-info {
  flex: 1;
  padding: 36px 32px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
}
.mf-name {
  font-size: 1.6rem;
  font-weight: 800;
  color: #111827;
  margin: 0 0 4px;
}
.mf-pos {
  font-size: .92rem;
  font-weight: 600;
  color: #2563eb;
  margin: 0 0 14px;
}
.mf-divider {
  width: 40px;
  height: 3px;
  background: #d4a017;
  border-radius: 2px;
  margin-bottom: 14px;
}
.mf-desc {
  font-size: .88rem;
  color: #6b7280;
  line-height: 1.7;
  margin: 0;
  white-space: pre-line;
}

/* ── REGULAR MODAL ── */
.modal-regular {
  max-width: 440px;
  width: 100%;
}
.mr-photo {
  width: 100%;
  height: 220px;
  background: #f3f4f6;
  position: relative;
  overflow: hidden;
}
.mr-photo img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  background: #f3f4f6;
}
.modal-photo-placeholder {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 3rem;
  color: #d1d5db;
}
.mr-info {
  padding: 20px 22px;
}
.mr-name {
  font-size: 1.2rem;
  font-weight: 800;
  color: #111827;
  margin: 0 0 4px;
}
.mr-pos {
  font-size: .85rem;
  font-weight: 600;
  color: #2563eb;
  margin: 0 0 12px;
}
.mr-desc {
  font-size: .85rem;
  color: #6b7280;
  line-height: 1.65;
  margin: 0;
  white-space: pre-line;
}

@media (max-width: 700px) {
  .modal-featured { flex-direction: column; height: auto; }
  .mf-photo { width: 100%; height: 240px; }
  .mf-info { padding: 24px 20px; }
}

/* ── EMPTY ── */
.empty-state {
  text-align: center; padding: 80px 20px; color: #999;
}
.empty-icon { font-size: 4rem; margin-bottom: 12px; }
.empty-text { font-size: 1rem; }

/* ── RESPONSIVE ── */
@media (max-width: 960px) {
  .hero-content { padding: 48px 40px 60px; }
  .line { font-size: 2.1rem; }
  .team-section { padding: 56px 40px 72px; }
  .featured-card { flex-direction: column; align-items: center; text-align: center; padding: 28px; gap: 16px; }
  .featured-photo { width: 100%; max-width: 280px; border-radius: 12px; aspect-ratio: 4/5; }
  .featured-photo img { position: relative; object-position: top center; }
  .featured-info { padding: 0; }
  .featured-btn { align-self: center; }
  .team-grid { grid-template-columns: repeat(2, 1fr); gap: 24px; }
  .orgchart-img-wrap { max-height: 500px; }
}
@media (max-width: 600px) {
  .line { font-size: 1.6rem; }
  .hero-content { padding: 36px 24px 52px; }
  .team-section { padding: 40px 20px 56px; }
  .section-title { font-size: 1.3rem; }
  .team-grid { grid-template-columns: 1fr; gap: 20px; max-width: 360px; margin: 0 auto; }
  .orgchart-img-wrap { max-height: 400px; }
}

/* ============ STRUKTUR ORGANISASI (gambar bagan) ============ */
.orgchart-section {
  padding: 0 40px 88px;
  background: #fff;
}
.orgchart-section .section-inner {
  max-width: 1500px;
}
.orgchart-img-wrap {
  width: 100%;
  margin: 0 auto;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 20px rgba(0,0,0,.08);
  max-height: 800px;
  overflow-y: hidden;
}
.orgchart-img {
  width: 100%;
  height: auto;
  display: block;
  min-height: 500px;
  object-fit: contain;
  object-position: top center;
}

@media (max-width: 960px) {
  .orgchart-section { padding: 0 24px 72px; }
}
@media (max-width: 600px) {
  .orgchart-section { padding: 0 16px 56px; }
}

/* ============ FOOTER ============ */
.site-footer {
  width: 100%;
  background-color: #5F9E42;
  padding: 25px 0;
  margin-top: 40px;
}
.footer-container {
  display: flex;
  align-items: center;
  justify-content: center;
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}
.footer-left-content { display: none; }
.footer-copyright { text-align: center; width: 100%; }
.footer-copyright p {
  font-family: Arial, sans-serif;
  font-size: 13px;
  color: #ffffff;
  margin: 0;
  font-weight: normal;
  letter-spacing: 0.3px;
}
</style>