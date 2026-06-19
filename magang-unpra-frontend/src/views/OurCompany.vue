<template>
  <div class="page">

    <!-- HERO -->
    <section class="hero" ref="heroEl">
      <div class="hero-bg">
        <img :src="profile.hero_image ? getImageUrl(profile.hero_image) : '/images/lokasi pabrik.jpeg'" alt="Pabrik TeL" class="hero-img" ref="heroImg" />
        <div class="hero-overlay"></div>
        <div class="hero-particles" ref="particlesEl"></div>
      </div>
      <div class="hero-content">
        <p class="eyebrow" :class="{in: h1}">ABOUT US</p>
        <h1 :class="{in: h1}">
          <span class="line" v-for="(l,i) in ['OUR COMPANY']" :key="i" :style="{transitionDelay: (0.18+i*0.14)+'s'}">{{l}}</span>
        </h1>
        <div class="accent-line" :class="{in: h1}" style="transition-delay:.52s"></div>
        <p v-if="profile.title" class="hero-subtitle" :class="{in: h1}" style="transition-delay:.54s">{{ profile.title }}</p>
        <p class="hero-p" :class="{in: h1}" style="transition-delay:.62s" v-html="profile.content || 'Memuat data...'"></p>
      </div>
      <div class="hero-scroll" :class="{in: h1}" style="transition-delay:.9s">
        <span></span>
      </div>
    </section>

    <!-- CREED -->
    <section class="creed" ref="creedEl">
      <!-- TAMBAHAN: background image -->
      <img :src="profile.creed_bg_image ? getImageUrl(profile.creed_bg_image) : '/images/backgroundour.jpeg'" alt="" class="creed-bg-img" aria-hidden="true" />
      <div class="creed-bg-kanji" aria-hidden="true">和</div>
      <div class="section-inner">
        <div class="section-eyebrow" :class="{in: c1}">
          <span class="line-l"></span>COMPANY CREED<span class="line-r"></span>
        </div>
        <div class="kanji-row">
          <div
            class="kanji-item"
            v-for="(k,i) in creedData"
            :key="i"
            :class="{in: c1}"
            :style="{transitionDelay: (0.1+i*0.18)+'s'}"
            @mouseenter="hoverKanji=i"
            @mouseleave="hoverKanji=-1"
          >
            <div class="kanji-char" :class="{hovered: hoverKanji===i}">{{k.title_jp || '?'}}</div>
            <div class="kanji-divider"></div>
            <p class="kanji-roma">{{k.roma || ''}}</p>
            <p class="kanji-name">{{k.title_en}}</p>
          </div>
        </div>
        <p class="creed-tagline" :class="{in: c1}" style="transition-delay:.7s">
          To achieve sustainable growth in the pulp and paper industry.
        </p>
      </div>
    </section>

    <!-- VALUES -->
    <section class="values" ref="valuesEl">
      <div class="section-inner">
        <div
          class="value-card"
          v-for="(v,i) in values"
          :key="i"
          :class="{in: v1}"
          :style="{transitionDelay: (i*0.15)+'s'}"
          @mouseenter="v.hovered=true"
          @mouseleave="v.hovered=false"
        >
          <div class="v-icon-wrap" :class="{hov: v.hovered}">
            <div class="v-icon" v-html="v.svg"></div>
          </div>
          <h3 class="v-title">{{v.title}}</h3>
          <p class="v-sub">{{v.tagline}}</p>
          <div class="v-bar" :class="{hov: v.hovered}"></div>
          <p class="v-desc">{{v.desc}}</p>
        </div>
      </div>
    </section>

    <!-- DOCUMENT -->
    <section class="doc" ref="docEl">
      <div class="section-inner doc-inner">
        <div class="doc-left" :class="{in: d1}">
          <div class="doc-card">
            <div class="doc-corner tl"></div>
            <div class="doc-corner tr"></div>
            <div class="doc-corner bl"></div>
            <div class="doc-corner br"></div>
            <img src="/images/pdf.png" alt="PDF Preview" class="doc-img" />
            <div class="doc-shine"></div>
          </div>
        </div>
        <div class="doc-right" :class="{in: d1}" style="transition-delay:.2s">
          <span class="doc-badge">{{ currentDoc.category || 'DOCUMENT PREVIEW' }}</span>
          <h2 class="doc-title">{{ currentDoc.title || 'Company Document' }}</h2>
          <div class="doc-meta-row">
            <span v-if="currentDoc.doc_date" class="doc-meta-item">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="4" width="18" height="18" rx="2"/><line x1="16" y1="2" x2="16" y2="6"/><line x1="8" y1="2" x2="8" y2="6"/><line x1="3" y1="10" x2="21" y2="10"/></svg>
              {{ currentDoc.doc_date }}
            </span>
            <span class="doc-meta-item">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
              {{ currentDoc.file_type || 'PDF' }} &nbsp;·&nbsp; {{ currentDoc.file_size || '-' }}
            </span>
          </div>
          <p class="doc-desc">{{ currentDoc.description || 'Document description.' }}</p>
          <div class="doc-btns">
            <a :href="docFileUrl(currentDoc)" download class="btn-dl">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="7 10 12 15 17 10"/><line x1="12" y1="15" x2="12" y2="3"/></svg>
              Download {{ currentDoc.file_type || 'PDF' }}
            </a>
            <button v-if="currentDoc.file_url" class="btn-view" @click="modal=true">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
              Quick View
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- MODAL -->
    <transition name="fade-modal">
      <div v-if="modal" class="modal-backdrop" @click.self="modal=false">
        <div class="modal-box" :class="{open: modal}">
          <button class="modal-close" @click="modal=false">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
          </button>
          <iframe :src="docFileUrl(currentDoc)" class="modal-frame"></iframe>
        </div>
      </div>
    </transition>

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
import { ref, onMounted, onUnmounted, computed } from 'vue'
import api from '@/services/api'

const BASE_URL = 'http://localhost:8080'

const heroEl = ref(null)
const creedEl = ref(null)
const valuesEl = ref(null)
const docEl = ref(null)
const heroImg = ref(null)
const particlesEl = ref(null)
const modal = ref(false)
const hoverKanji = ref(-1)

const h1 = ref(false)
const c1 = ref(false)
const v1 = ref(false)
const d1 = ref(false)

// ── DATA DENGAN FALLBACK DEFAULT ──
// Data ini akan ditimpa oleh API jika database sudah terisi

const defaultProfileText = `PT Tanjungenim Lestari Pulp and Paper ( TeL) is world class manufacturer of high product quality and environmental friendly market pulp mill. This was established on June 18, 1990, commenced construction in mid-1997 and the commercial operation started on May, 2000 . The mill is located in 1,250 ha area in the Banuayu village, District Empat Petulai Dangku, Muara Enim Regency, South Sumatra province, Indonesia.

TeL is a Foreign Investment Company (PMA)- Marubeni Corporation , Japan , as the National Vital Objects Industrial sector (OVNI) declared by the Minister of Industry in 2014 . The mill has market pulp production capacity of 490,000 Adt / year. Presently mill has 1600 employees and support workforce together where ~ 80% of them are residents of South Sumatra`

const profile = ref({
  title: 'PT Tanjungenim Lestari Pulp and Paper',
  content: defaultProfileText,
  address: '',
  phone: '',
  email: '',
  hero_image: '',
  creed_bg_image: '',
})

const defaultCreeds = [
  { id: null, title_jp: '和', title_en: 'Harmony', roma: 'WA', tagline: 'To respect each other and cooperate.', description: 'We shall stay in touch with society and stakeholders by engaging in corporate activities which advance our credibility as preferred principals.', sort_order: 0, is_active: true },
  { id: null, title_jp: '新', title_en: 'Innovation', roma: 'SHIN', tagline: 'To be active and innovative.', description: 'We shall constantly strive hard to improve our performance.', sort_order: 1, is_active: true },
  { id: null, title_jp: '心', title_en: 'Heart', roma: 'SHIN', tagline: 'To be fair and decent.', description: 'We shall comply with the laws and follow fair corporate practices.', sort_order: 2, is_active: true },
]

const creedData = ref([...defaultCreeds])

const defaultDoc = {
  id: null,
  title: 'PEFC Chain of Custody Commitment Statement',
  category: 'DOCUMENT PREVIEW',
  doc_date: 'April 2023',
  file_type: 'PDF',
  file_size: '1.2 MB',
  description: 'PT. Tanjungenim Lestari Pulp and Paper is committed to sourcing wood-based products in an environmentally and socially responsible manner, in full compliance with EU Timber Regulation and sustainable forest management practices.',
  file_url: '/files/PEFC_Statement.pdf',
  _is_fallback: true,
}

const documents = ref([{ ...defaultDoc }])

const currentDoc = computed(() =>
  documents.value.length > 0 ? documents.value[0] : defaultDoc
)

function docFileUrl(doc) {
  if (!doc || !doc.file_url) { return '/files/PEFC_Statement.pdf' }
  if (doc._is_fallback) { return doc.file_url }
  return BASE_URL + doc.file_url
}

function getImageUrl(path) {
  if (!path) { return '' }
  if (path.startsWith('http')) { return path }
  if (path.startsWith('/uploads/')) { return BASE_URL + path }
  if (path.startsWith('uploads/')) { return BASE_URL + '/' + path }
  return path
}

const defaultSvgs = [
  `<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6"><path d="M12 3v18"/><path d="M3 9h18"/><path d="M5.5 9L3 16a4.5 4.5 0 009 0L9.5 9"/><path d="M14.5 9L12 16a4.5 4.5 0 009 0L18.5 9"/></svg>`,
  `<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6"><circle cx="12" cy="10" r="5"/><path d="M12 15v4"/><path d="M9 19h6"/><path d="M9.5 10a2.5 2.5 0 012.5-2.5"/></svg>`,
  `<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6"><circle cx="9" cy="7" r="3"/><circle cx="15" cy="7" r="3"/><path d="M3 21v-1a6 6 0 016-6h6a6 6 0 016 6v1"/></svg>`,
  `<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6"><path d="M22 12h-4l-3 9L9 3l-3 9H2"/></svg>`,
  `<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6"><circle cx="12" cy="8" r="5"/><path d="M3 21v-2a7 7 0 0114 0v2"/></svg>`,
]

const values = computed(() =>
  creedData.value.map((c, i) => ({
    title: c.title_en,
    tagline: c.tagline,
    desc: c.description,
    hovered: false,
    svg: defaultSvgs[i % defaultSvgs.length],
  }))
)

function makeParticles() {
  if (!particlesEl.value) {return}
  const el = particlesEl.value
  for (let i = 0; i < 18; i++) {
    const p = document.createElement('span')
    p.className = 'particle'
    p.style.cssText = `
      left:${Math.random()*100}%;
      top:${Math.random()*100}%;
      width:${2+Math.random()*3}px;
      height:${2+Math.random()*3}px;
      animation-delay:${Math.random()*6}s;
      animation-duration:${4+Math.random()*6}s;
    `
    el.appendChild(p)
  }
}

function useObserver(el, flag, threshold = 0.2) {
  const io = new IntersectionObserver(([e]) => {
    if (e.isIntersecting) { flag.value = true; io.disconnect() }
  }, { threshold })
  if (el) {io.observe(el)}
  return io
}

const observers = []
onMounted(async () => {
  makeParticles()
  setTimeout(() => { h1.value = true }, 120)
  observers.push(useObserver(creedEl.value, c1))
  observers.push(useObserver(valuesEl.value, v1))
  observers.push(useObserver(docEl.value, d1))

  try {
    const [profileRes, creedRes, docRes] = await Promise.all([
      api.get('/company-profile'),
      api.get('/creeds'),
      api.get('/company-documents'),
    ])
    // Profile: timpa hanya jika API mengembalikan data dengan konten
    const pData = profileRes.data?.data || profileRes.data
    if (pData && (pData.content || pData.title)) {
      profile.value = {
        title: pData.title || profile.value.title,
        content: pData.content || profile.value.content,
        hero_image: pData.hero_image || '',
        creed_bg_image: pData.creed_bg_image || '',
      }
    }
    // Creed: timpa hanya jika ada data aktif dari API
    const allCreeds = Array.isArray(creedRes.data) ? creedRes.data : (creedRes.data.data || [])
    const activeCreeds = allCreeds.filter(c => c.is_active)
    if (activeCreeds.length > 0) {
      creedData.value = activeCreeds
    }
    // Dokumen: timpa hanya jika ada data dari API
    const apiDocs = Array.isArray(docRes.data) ? docRes.data : (docRes.data.data || [])
    if (apiDocs.length > 0) {
      documents.value = apiDocs
    }
  } catch (e) {
    console.error('Gagal memuat data Our Company, menggunakan data default:', e)
  }
})
onUnmounted(() => observers.forEach(o => o.disconnect()))
</script>

<style scoped>
*{box-sizing:border-box;margin:0;padding:0}

.page{
  font-family:'Segoe UI',system-ui,Arial,sans-serif;
  color:#1a1a1a;
  background:#fff;
  overflow-x:hidden;
}

/* ── HERO ── */
.hero{
  position:relative;
  min-height:520px;
  display:flex;
  align-items:flex-end;
  overflow:hidden;
}
.hero-bg{position:absolute;inset:0;z-index:0}
.hero-img{
  width:100%;height:100%;object-fit:cover;
  object-position:center 40%;
  display:block;
  transform:scale(1.06);
  animation:heroZoom 14s ease-out forwards;
}
@keyframes heroZoom{to{transform:scale(1)}}

.hero-overlay{
  position:absolute;inset:0;
  background:linear-gradient(105deg,rgba(5,28,14,.88) 0%,rgba(5,28,14,.6) 42%,rgba(5,28,14,.18) 100%);
}
.hero-particles{position:absolute;inset:0;pointer-events:none}

.hero-content{
  position:relative;z-index:2;
  padding:64px 72px 72px;
  max-width:660px;
}

.eyebrow{
  font-size:.7rem;font-weight:700;letter-spacing:.16em;
  color:#5ecb7a;margin-bottom:18px;
  opacity:0;transform:translateY(14px);
  transition:opacity .6s,transform .6s;
}
.eyebrow.in{opacity:1;transform:none}

h1{margin:0 0 22px;color:#fff}
.line{
  display:block;
  font-size:2.8rem;font-weight:800;line-height:1.18;
  opacity:0;transform:translateY(22px) skewY(1.5deg);
  transition:opacity .7s cubic-bezier(.22,1,.36,1),transform .7s cubic-bezier(.22,1,.36,1);
}
.in .line{opacity:1;transform:none}

.accent-line{
  width:0;height:3px;background:#2d7a3d;
  border-radius:2px;margin-bottom:22px;
  transition:width .8s cubic-bezier(.22,1,.36,1);
}
.accent-line.in{width:56px}

.hero-subtitle{
  font-size:.82rem;color:#5ecb7a;font-weight:600;margin-bottom:8px;
  opacity:0;transform:translateY(14px);
  transition:opacity .7s,transform .7s;
}
.hero-subtitle.in{opacity:1;transform:none}
.hero-p{
  font-size:.93rem;color:rgba(255,255,255,.8);line-height:1.8;
  opacity:0;transform:translateY(14px);
  transition:opacity .7s,transform .7s;
}
.hero-p.in{opacity:1;transform:none}

.hero-scroll{
  position:absolute;bottom:28px;left:50%;transform:translateX(-50%);
  z-index:2;display:flex;flex-direction:column;align-items:center;
  opacity:0;transition:opacity 1s;
}
.hero-scroll.in{opacity:1}
.hero-scroll span{
  display:block;width:1.5px;height:36px;
  background:linear-gradient(to bottom,rgba(255,255,255,.5),transparent);
  animation:scrollPulse 2s infinite;
}
@keyframes scrollPulse{0%,100%{opacity:.3;transform:scaleY(.4) translateY(-6px)}50%{opacity:1;transform:scaleY(1) translateY(0)}}

/* particles */
:deep(.particle){
  position:absolute;background:rgba(255,255,255,.18);
  border-radius:50%;
  animation:floatUp linear infinite;
}
@keyframes floatUp{0%{opacity:0;transform:translateY(0) scale(.5)}20%{opacity:1}80%{opacity:.4}100%{opacity:0;transform:translateY(-120px) scale(1.2)}}

/* ── CREED ── */
.creed{
  position:relative;
  background:#f8f8f6;
  padding:96px 72px;
  text-align:center;
  overflow:hidden;
}

/* TAMBAHAN: background image */
.creed-bg-img{
  position:absolute;
  inset:0;
  width:100%;
  height:100%;
  object-fit:cover;
  object-position:center;
  pointer-events:none;
  user-select:none;
  opacity:0.45;
  z-index:0;
}

.creed-bg-kanji{
  position:absolute;left:-40px;top:50%;transform:translateY(-50%);
  font-size:22rem;color:rgba(45,122,61,.04);
  font-family:serif;line-height:1;pointer-events:none;
  user-select:none;
  z-index:1;
}
.section-inner{max-width:1100px;margin:0 auto;position:relative;z-index:2}

.section-eyebrow{
  display:flex;align-items:center;justify-content:center;
  gap:18px;font-size:.68rem;font-weight:700;
  letter-spacing:.16em;color:#2d7a3d;margin-bottom:56px;
  opacity:0;transform:translateY(16px);
  transition:opacity .7s,transform .7s;
}
.section-eyebrow.in{opacity:1;transform:none}
.line-l,.line-r{flex:1;max-width:100px;height:1px;background:#ccc}

.kanji-row{
  display:flex;justify-content:center;
  gap:0;margin-bottom:40px;
}
.kanji-item{
  flex:1;max-width:220px;padding:0 24px 32px;
  position:relative;cursor:default;
  opacity:0;transform:translateY(32px);
  transition:opacity .7s cubic-bezier(.22,1,.36,1),transform .7s cubic-bezier(.22,1,.36,1);
}
.kanji-item.in{opacity:1;transform:none}
.kanji-item:not(:last-child)::after{
  content:'';position:absolute;right:0;top:20%;bottom:20%;
  width:1px;background:#ddd;
}
.kanji-char{
  font-size:5.5rem;color:#1a3f28;font-family:serif;line-height:1;
  margin-bottom:16px;
  transition:transform .4s cubic-bezier(.34,1.56,.64,1),color .3s;
}
.kanji-char.hovered{transform:scale(1.18) translateY(-6px);color:#2d7a3d}
.kanji-divider{width:28px;height:2px;background:#2d7a3d;margin:0 auto 14px;border-radius:1px}
.kanji-roma{font-size:.64rem;letter-spacing:.18em;color:#aaa;margin-bottom:6px}
.kanji-name{font-size:.88rem;font-weight:700;color:#444}
.creed-tagline{
  font-size:.98rem;color:#666;font-style:italic;
  opacity:0;transform:translateY(14px);
  transition:opacity .7s,transform .7s;
}
.creed-tagline.in{opacity:1;transform:none}

/* ── VALUES ── */
.values{padding:88px 72px;background:#fff}
.values .section-inner{
  display:grid;
  grid-template-columns:repeat(3,1fr);
  gap:28px;
}
.value-card{
  border:1px solid #e8e8e8;border-radius:8px;
  padding:36px 28px;
  opacity:0;transform:translateY(36px);
  transition:opacity .75s cubic-bezier(.22,1,.36,1),transform .75s cubic-bezier(.22,1,.36,1),box-shadow .35s,border-color .35s;
  position:relative;overflow:hidden;
}
.value-card::before{
  content:'';position:absolute;inset:0;
  background:linear-gradient(135deg,rgba(45,122,61,.04) 0%,transparent 60%);
  opacity:0;transition:opacity .4s;
}
.value-card:hover::before{opacity:1}
.value-card:hover{box-shadow:0 16px 48px rgba(45,122,61,.12);border-color:#c8e6d0}
.value-card.in{opacity:1;transform:none}

.v-icon-wrap{
  width:60px;height:60px;border-radius:50%;
  background:#edf7f0;display:flex;align-items:center;justify-content:center;
  margin-bottom:20px;color:#2d7a3d;
  transition:transform .4s cubic-bezier(.34,1.56,.64,1),background .3s;
}
.v-icon-wrap.hov{transform:scale(1.15) rotate(-6deg);background:#d4edda}
.v-title{font-size:1.1rem;font-weight:700;color:#1a1a1a;margin-bottom:6px}
.v-sub{font-size:.82rem;color:#2d7a3d;font-style:italic;margin-bottom:14px}
.v-bar{width:28px;height:2px;background:#2d7a3d;border-radius:1px;margin-bottom:14px;transition:width .4s}
.v-bar.hov{width:48px}
.v-desc{font-size:.86rem;color:#666;line-height:1.75}

/* ── DOC ── */
.doc{padding:88px 72px;background:#f8f8f6}
.doc-inner{
  display:grid;grid-template-columns:280px 1fr;
  gap:64px;align-items:center;
}
.doc-left,.doc-right{
  opacity:0;transform:translateY(28px);
  transition:opacity .8s cubic-bezier(.22,1,.36,1),transform .8s cubic-bezier(.22,1,.36,1);
}
.doc-left.in,.doc-right.in{opacity:1;transform:none}

.doc-card{
  position:relative;border-radius:8px;overflow:hidden;
  box-shadow:0 8px 40px rgba(0,0,0,.14);
}
.doc-img{width:100%;display:block}
.doc-shine{
  position:absolute;inset:0;
  background:linear-gradient(135deg,rgba(255,255,255,.18) 0%,transparent 55%);
  pointer-events:none;
}
.doc-corner{
  position:absolute;width:18px;height:18px;
  border-color:#2d7a3d;border-style:solid;
  z-index:2;
}
.doc-corner.tl{top:8px;left:8px;border-width:2px 0 0 2px;border-radius:3px 0 0 0}
.doc-corner.tr{top:8px;right:8px;border-width:2px 2px 0 0;border-radius:0 3px 0 0}
.doc-corner.bl{bottom:8px;left:8px;border-width:0 0 2px 2px;border-radius:0 0 0 3px}
.doc-corner.br{bottom:8px;right:8px;border-width:0 2px 2px 0;border-radius:0 0 3px 0}

.doc-badge{
  font-size:.65rem;font-weight:700;letter-spacing:.12em;
  color:#2d7a3d;border:1.5px solid #2d7a3d;
  padding:3px 10px;border-radius:3px;
  display:inline-block;margin-bottom:16px;
}
.doc-title{font-size:1.35rem;font-weight:800;line-height:1.38;color:#111;margin-bottom:16px}
.doc-meta-row{display:flex;gap:20px;margin-bottom:18px}
.doc-meta-item{display:flex;align-items:center;gap:5px;font-size:.78rem;color:#888}
.doc-desc{font-size:.875rem;color:#555;line-height:1.78;margin-bottom:28px}
.doc-btns{display:flex;gap:12px}

.btn-dl{
  display:inline-flex;align-items:center;gap:7px;
  padding:12px 24px;background:#2d7a3d;color:#fff;
  border:none;border-radius:5px;font-size:.88rem;font-weight:600;
  cursor:pointer;text-decoration:none;
  transition:background .2s,transform .2s;
}
.btn-dl:hover{background:#1a5c2a;transform:translateY(-1px)}
.btn-view{
  display:inline-flex;align-items:center;gap:7px;
  padding:12px 24px;background:transparent;
  color:#2d7a3d;border:1.5px solid #2d7a3d;
  border-radius:5px;font-size:.88rem;font-weight:600;
  cursor:pointer;transition:all .2s;
}
.btn-view:hover{background:#eaf5ee;transform:translateY(-1px)}

/* ── MODAL ── */
.modal-backdrop{
  position:fixed;inset:0;z-index:1000;
  background:rgba(0,0,0,.65);
  display:flex;align-items:center;justify-content:center;
}
.modal-box{
  background:#fff;border-radius:10px;
  width:82vw;height:87vh;position:relative;
  overflow:hidden;
  transform:scale(.92);opacity:0;
  transition:transform .35s cubic-bezier(.34,1.56,.64,1),opacity .3s;
}
.modal-box.open{transform:scale(1);opacity:1}
.modal-close{
  position:absolute;top:12px;right:12px;z-index:10;
  background:#fff;border:1px solid #ddd;
  border-radius:50%;width:34px;height:34px;
  cursor:pointer;display:flex;align-items:center;justify-content:center;
  transition:background .2s;
}
.modal-close:hover{background:#fee}
.modal-frame{width:100%;height:100%;border:none}
.fade-modal-enter-active,.fade-modal-leave-active{transition:opacity .25s}
.fade-modal-enter-from,.fade-modal-leave-to{opacity:0}

/* ── RESPONSIVE ── */
@media(max-width:960px){
  .hero-content{padding:48px 40px 60px}
  .line{font-size:2.1rem}
  .creed{padding:72px 40px}
  .values{padding:72px 40px}
  .values .section-inner{grid-template-columns:1fr}
  .doc{padding:72px 40px}
  .doc-inner{grid-template-columns:1fr;gap:40px}
}
@media(max-width:600px){
  .line{font-size:1.6rem}
  .hero-content{padding:36px 24px 52px}
  .kanji-char{font-size:3.8rem}
  .doc-btns{flex-direction:column}
}

/* ============ FOOTER ============ */
.site-footer {
  width: 100%;
  /* Warna hijau tua disesuaikan dengan identitas PT TEL pada gambar */
  background-color: #5F9E42; 
  padding: 25px 0; /* Memberikan ruang tinggi baris yang pas */
  margin-top: 40px; /* Jarak pemisah dari konten di atasnya */
}

.footer-container {
  display: flex;
  align-items: center;
  justify-content: center; /* DIUBAH: Membuat konten berkumpul di tengah */
  max-width: 1140px; /* Menyelaraskan lebar dengan konten atas website */
  margin: 0 auto;
  padding: 0 40px; /* Jarak aman agar teks tidak terlalu menempel ke pinggir layar */
}

/* Kolom Kiri Footer */
.footer-left-content {
  display: none; /* DIUBAH: Disembunyikan karena kosong, agar tidak memakan space */
}

/* Kolom Kanan Footer (Teks Copyright) */
.footer-copyright {
  text-align: center; /* DIUBAH: Memastikan teks rata tengah */
  width: 100%; /* Memastikan block mengambil ruang penuh untuk fungsi centering */
}

.footer-copyright p {
  font-family: Arial, sans-serif;
  font-size: 13px; /* Ukuran teks kecil, tipis, dan normal untuk footer */
  color: #ffffff;  /* Warna teks putih bersih agar kontras dengan background hijau */
  margin: 0;
  font-weight: normal;
  letter-spacing: 0.3px;
}
</style>