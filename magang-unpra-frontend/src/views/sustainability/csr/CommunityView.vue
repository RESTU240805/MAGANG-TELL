<template>
  <div class="community-page">

    <!-- ─── Hero ─────────────────────────────────────────────── -->
    <section class="hero">
      <div class="hero-bg"></div>
      <div class="hero-content">
        <span class="hero-eyebrow">SUSTAINABILITY & COMMUNITY</span>
        <h1 class="hero-title">Local Community<br />Development</h1>
        <p class="hero-desc">
          We believe that business sustainability goes hand in hand with the welfare
          of local communities. Through structured CSR initiatives, we build a
          better future together.
        </p>
        
      </div>
    </section>

    <!-- ─── Local Community Development ─────────────────────── -->
    <section class="section-lcd" ref="lcdRef" :class="{'reveal-in': lcdVisible}">
      <div class="lcd-card">
        <!-- Left: circle illustration -->
        <div class="lcd-left">
          <div class="lcd-circle">
            <svg width="56" height="56" viewBox="0 0 24 24" fill="none" stroke="#16a34a" stroke-width="1.2">
              <path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2"/>
              <circle cx="9" cy="7" r="4"/>
              <path d="M23 21v-2a4 4 0 00-3-3.87"/>
              <path d="M16 3.13a4 4 0 010 7.75"/>
            </svg>
            <p class="lcd-circle-text">
              Strategi Tanggung Jawab Sosial Perusahaan kami didasarkan pada pemberdayaan
              berkelanjutan, bukan sekadar filantropi. Kami berfokus pada dua pilar utama pengembangan:
            </p>
          </div>
          <div class="lcd-leaf">🌿</div>
        </div>

        <!-- Right: programs -->
        <div class="lcd-right">
          <div class="lcd-title-row">
            <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="#16a34a" stroke-width="2">
              <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z"/>
              <path d="M8 12l2 2 4-4"/>
            </svg>
            <h2>Local Community Development</h2>
          </div>
          <div class="lcd-divider"></div>

          <div class="lcd-program" v-for="(p, i) in programs" :key="i"
            :style="{ transitionDelay: lcdVisible ? `${i * 150 + 200}ms` : '0ms' }"
            :class="{'lcd-program--visible': lcdVisible}">
            <div class="lcd-program-icon">
              <!-- icon community -->
              <svg v-if="i===0" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="#16a34a" stroke-width="1.5">
                <path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2"/>
                <circle cx="9" cy="7" r="4"/>
                <path d="M23 21v-2a4 4 0 00-3-3.87"/>
                <path d="M16 3.13a4 4 0 010 7.75"/>
              </svg>
              <!-- icon relations -->
              <svg v-else width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="#16a34a" stroke-width="1.5">
                <circle cx="12" cy="12" r="10"/>
                <path d="M8.56 2.75c4.37 6.03 6.02 9.42 8.03 17.72m2.54-15.38c-3.72 4.35-8.94 5.66-16.88 5.85m19.5 1.9c-3.5-.93-6.63-.82-8.94 0-2.58.92-5.01 2.86-7.44 6.32"/>
              </svg>
            </div>
            <div>
              <h3 class="lcd-program-title">{{ p.title }}</h3>
              <p class="lcd-program-desc">{{ p.desc }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ─── CSR Target ────────────────────────────────────────── -->
    <section class="section-target" ref="targetRef" :class="{'reveal-in': targetVisible}">
      <div class="target-inner">
        <div class="target-left">
          <h2 class="target-title">CSR Target</h2>
          <div class="target-divider"></div>
          <p class="target-subtitle">The objectives of the CSR Program are:</p>
          <div class="target-grid">
            <div class="target-item" v-for="(item, i) in csrTargets" :key="i"
              :style="{ transitionDelay: targetVisible ? `${i * 100 + 150}ms` : '0ms' }"
              :class="{'target-item--visible': targetVisible}">
              <div class="target-icon">
                <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="#16a34a" stroke-width="1.5">
                  <component :is="'path'" :d="targetIcons[i]" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </div>
              <div>
                <p class="target-item-text">{{ item.main }}</p>
                <p v-if="item.sub" class="target-item-sub">{{ item.sub }}</p>
              </div>
            </div>
          </div>
        </div>
        <div class="target-right">
          <img src="/images/csr-economy-1.jpeg" alt="CSR Activity"
            :class="{'img-reveal': targetVisible}"
            class="target-img"
            @error="(e) => e.target.src='https://placehold.co/400x500/dcfce7/16a34a?text=CSR'" />
        </div>
      </div>
    </section>

    <!-- ─── Main Program Pillars ──────────────────────────────── -->
    <section class="section-pilar">
      <div class="container">
        <div class="pilar-header" ref="pilarRef" :class="{'reveal-in': pilarVisible}">
          <h2>Main Program Pillars
            <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="#16a34a" stroke-width="2" style="display:inline;margin-left:6px;vertical-align:middle">
              <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z"/>
              <path d="M8 12l2 2 4-4"/>
            </svg>
          </h2>
          <p>Click on the categories below to view the details of our activities and real contributions to the community.</p>
        </div>

        <!-- Tab Nav -->
        <div class="tab-nav">
          <button v-for="(tab, i) in tabs" :key="i"
            @click="selectTab(i)"
            :class="['tab-btn', { active: activeTab === i }]">
            <span class="tab-icon" v-html="tab.icon"></span>
            {{ tab.label }}
          </button>
        </div>

        <!-- Tab Content -->
        <Transition name="tab-fade" mode="out-in">
          <div class="tab-content" :key="activeTab">
            <div class="tab-left">
              <div class="tab-image-wrap">
                <img :src="tabs[activeTab].images[activeImage]" :alt="tabs[activeTab].label"
                  class="tab-image"
                  @error="(e) => e.target.src='https://placehold.co/500x380/dcfce7/16a34a?text='+tabs[activeTab].label" />
              </div>

              <!-- Thumbnail gallery, only shown when a tab has more than one photo -->
              <div class="tab-thumbs" v-if="tabs[activeTab].images.length > 1">
                <button v-for="(img, gi) in tabs[activeTab].images" :key="gi"
                  class="tab-thumb-btn"
                  :class="{ 'tab-thumb-btn--active': activeImage === gi }"
                  @click="activeImage = gi">
                  <img :src="img" :alt="`${tabs[activeTab].label} ${gi + 1}`" />
                </button>
              </div>
            </div>
            <div class="tab-right">
              <h3 class="tab-title">{{ tabs[activeTab].title }}</h3>
              <div class="tab-divider"></div>
              <p class="tab-desc">{{ tabs[activeTab].desc }}</p>
              <div class="tab-quote">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#16a34a" stroke-width="1.5" style="flex-shrink:0;margin-top:2px">
                  <path d="M3 21c3 0 7-1 7-8V5c0-1.25-.756-2.017-2-2H4c-1.25 0-2 .75-2 1.972V11c0 1.25.75 2 2 2 1 0 1 0 1 1v1c0 1-1 2-2 2s-1 .008-1 1.031V20c0 1 0 1 1 1z"/>
                  <path d="M15 21c3 0 7-1 7-8V5c0-1.25-.757-2.017-2-2h-4c-1.25 0-2 .75-2 1.972V11c0 1.25.75 2 2 2h.75c0 2.25.25 4-2.75 4v3c0 1 0 1 1 1z"/>
                </svg>
                <p>{{ tabs[activeTab].quote }}</p>
              </div>
            </div>
          </div>
        </Transition>
      </div>
    </section>

    <!-- ─── CTA ───────────────────────────────────────────────── -->
    <template>
  <div class="community-page">

    <footer class="final-footer">
      <div class="footer-container">
        <p>Copyright 2019 PT TELPP. All right reserved.</p>
      </div>
    </footer>

  </div>
</template>

  </div>
  
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { RouterLink } from 'vue-router'

const activeTab = ref(0)
const activeImage = ref(0)

function selectTab(i) {
  activeTab.value = i
  activeImage.value = 0
}

// ─── Scroll reveal refs ──────────────────────────────────────
const lcdRef    = ref(null)
const targetRef = ref(null)
const pilarRef  = ref(null)
const ctaRef    = ref(null)
const tabsRef   = ref(null)

const lcdVisible    = ref(false)
const targetVisible = ref(false)
const pilarVisible  = ref(false)
const ctaVisible    = ref(false)
const tabsVisible   = ref(false)

let observer
onMounted(() => {
  observer = new IntersectionObserver((entries) => {
    entries.forEach(e => {
      if (!e.isIntersecting) return
      if (e.target === lcdRef.value)    lcdVisible.value    = true
      if (e.target === targetRef.value) targetVisible.value = true
      if (e.target === pilarRef.value)  pilarVisible.value  = true
      if (e.target === ctaRef.value)    ctaVisible.value    = true
      if (e.target === tabsRef.value)   tabsVisible.value   = true
    })
  }, { threshold: 0.12 })
  ;[lcdRef, targetRef, pilarRef, ctaRef, tabsRef].forEach(r => {
    if (r.value) observer.observe(r.value)
  })
})
onUnmounted(() => observer?.disconnect())

const programs = [
  {
    title: 'Community Development Program',
    desc: 'The Company\'s Corporate Social Responsibility strategy is based on sustainable empowerment, not mere philanthropy. We focus on two main pillars of development.',
  },
  {
    title: 'Relationship / Public Relations Development Program (Community Relations).',
    desc: 'The Company\'s Corporate Social Responsibility strategy is based on sustainable empowerment, not mere philanthropy. We focus on two main pillars of development.',
  },
]

const csrTargets = [
  { main: 'Empowerment of local human resources', sub: '(students, youth and students included)' },
  { main: 'Economic Empowerment', sub: 'of local communities around the company' },
  { main: 'Construction of social / public facilities', sub: '' },
  { main: 'Public health development', sub: '' },
  { main: 'Community and religious education', sub: '' },
  { main: 'Social Culture etc.', sub: '' },
]

const targetIcons = [
  'M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2M9 7a4 4 0 100 8 4 4 0 000-8zM23 21v-2a4 4 0 00-3-3.87M16 3.13a4 4 0 010 7.75',
  'M3 9l9-7 9 7v11a2 2 0 01-2 2H5a2 2 0 01-2-2zM9 22V12h6v10',
  'M3 21h18M9 3h6M10 3v18M14 3v18',
  'M20.84 4.61a5.5 5.5 0 00-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 00-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 000-7.78z',
  'M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z',
  'M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0118 0z',
]

const tabIconSvg = (d) => `<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="${d}" stroke-linecap="round" stroke-linejoin="round"/></svg>`

// Each tab now carries an `images` array (the real CSR documentation
// photos) instead of a single placeholder `image`. The first photo in
// the array is shown by default; the gallery thumbnails below let the
// visitor browse the rest.
const tabs = [
  {
    label: 'Education',
    icon: tabIconSvg('M12 14l9-5-9-5-9 5 9 5zM12 14l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14z'),
    title: 'Education',
    desc: 'Prepare a generation of quality sources from internal employees as well as from the village community around the company. We achieve this through scholarships, traineeships, Community Education Development, school foundation programs and school infrastructure improvements.',
    quote: '"We do not just provide scholarships, we build a generation ready to compete globally."',
    images: [
      '/images/csr-education-1.jpg.jpeg',
      '/images/csr-education-2.jpg.jpeg',
    ],
  },
  {
    label: 'Local Economy Development',
    icon: tabIconSvg('M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2M9 7a4 4 0 100 8 4 4 0 000-8zM23 21v-2a4 4 0 00-3-3.87M16 3.13a4 4 0 010 7.75'),
    title: 'Local Economy Development',
    desc: 'Growing and strengthening community businesses with short, medium, and long-term impact to be able to compete competitively in the business world in order to achieve equitable and sustainable community welfare around the company.',
    quote: '"We do not just provide assistance, we build economic independence through intensive mentoring for local SMEs."',
    images: [
      '/images/csr-economy-1.jpeg',
      '/images/csr-economy-2.jpg.jpeg',
    ],
  },
  {
    label: 'Healthy',
    icon: tabIconSvg('M20.84 4.61a5.5 5.5 0 00-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 00-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 000-7.78z'),
    title: 'Healthy',
    desc: 'Improving the health status of communities around the company through sustainable preventive, curative, and promotive health programs to create a healthy and productive society.',
    quote: '"Community health is a long-term investment that we prioritize together."',
    images: [
      '/images/csr-healthy-1.png',
      '/images/csr-healthy-2.png',
      '/images/csr-healthy-3.png',
      '/images/csr-healthy-4.png',
      '/images/csr-healthy-5.png',
    ],
  },
  {
    label: 'Infrastructure & Social',
    icon: tabIconSvg('M3 9l9-7 9 7v11a2 2 0 01-2 2H5a2 2 0 01-2-2zM9 22V12h6v10'),
    title: 'Infrastructure & Social Community',
    desc: 'Building and strengthening social infrastructure and public facilities that support the lives of communities around the company, including the construction of places of worship, sports facilities, and other public facilities.',
    quote: '"Strong infrastructure is the foundation of a prosperous and harmonious community."',
    images: [
      '/images/csr-social-1.png',
      '/images/csr-social-2.png',
      '/images/csr-social-3.png',
      '/images/csr-social-4.png',
    ],
  },
]
</script>

<style scoped>
* { box-sizing: border-box; }
.community-page { font-family: 'Segoe UI', Arial, sans-serif; color: #1a1a1a; background: #fff; }
.container { max-width: 1000px; margin: 0 auto; padding: 0 32px; }

/* ─── Hero ──────────────────────────────────────────────────── */
.hero {
  position: relative;
  min-height: 420px;
  display: flex;
  align-items: flex-end;
  overflow: hidden;
}
.hero-bg {
  position: absolute; inset: 0;
  background: linear-gradient(to right, rgba(0,0,0,0.75) 0%, rgba(0,0,0,0.15) 65%),
    url('/images/gedung.jpeg') center/cover no-repeat;
  background-color: #1a3a2a;
  animation: heroZoom 7s ease-out forwards;
}
@keyframes heroZoom {
  from { transform: scale(1.05); }
  to   { transform: scale(1); }
}
.hero-content {
  position: relative; z-index: 2;
  padding: 0 0 60px 72px;
  max-width: 500px;
  animation: fadeUp 0.75s ease both;
}
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(24px); }
  to   { opacity: 1; transform: translateY(0); }
}
.hero-eyebrow {
  display: block; font-size: 0.68rem; font-weight: 700;
  letter-spacing: 0.15em; color: #4ade80;
  text-transform: uppercase; margin-bottom: 12px;
}
.hero-title {
  font-size: 3rem; font-weight: 900; color: #fff;
  line-height: 1.08; margin: 0 0 16px;
}
.hero-desc {
  font-size: 0.88rem; color: rgba(255,255,255,0.8);
  line-height: 1.7; margin: 0 0 28px;
}
.hero-btn {
  display: inline-flex; align-items: center; gap: 8px;
  background: #16a34a; color: #fff;
  padding: 11px 22px; border-radius: 8px;
  font-size: 0.85rem; font-weight: 600;
  text-decoration: none;
  transition: background 0.2s, transform 0.2s;
}
.hero-btn:hover { background: #15803d; transform: translateY(-1px); }

/* ─── LCD Section ───────────────────────────────────────────── */
.section-lcd { padding: 60px 32px; background: #f9fafb; }
.lcd-card {
  max-width: 1000px; margin: 0 auto;
  background: #fff; border-radius: 20px;
  box-shadow: 0 4px 32px rgba(0,0,0,0.07);
  display: grid; grid-template-columns: 280px 1fr;
  overflow: hidden;
}
.lcd-left {
  background: #f0fdf4;
  padding: 40px 28px;
  display: flex; flex-direction: column;
  align-items: center; gap: 24px;
  position: relative;
}
.lcd-circle {
  width: 200px; height: 200px;
  background: #fff;
  border-radius: 50%;
  display: flex; flex-direction: column;
  align-items: center; justify-content: center;
  gap: 12px;
  padding: 24px;
  box-shadow: 0 2px 12px rgba(22,163,74,0.12);
}
.lcd-circle-text {
  font-size: 0.68rem; color: #555;
  line-height: 1.5; text-align: center; margin: 0;
}
.lcd-leaf { font-size: 2rem; }

.lcd-right { padding: 40px 40px; }
.lcd-title-row {
  display: flex; align-items: center; gap: 10px;
  margin-bottom: 8px;
}
.lcd-title-row h2 { font-size: 1.4rem; font-weight: 900; margin: 0; color: #111; }
.lcd-divider { width: 48px; height: 3px; background: #16a34a; border-radius: 2px; margin-bottom: 28px; }

.lcd-program {
  display: flex; align-items: flex-start; gap: 16px;
  padding: 18px 0;
  border-bottom: 1px solid #f3f4f6;
}
.lcd-program:last-child { border-bottom: none; }
.lcd-program-icon {
  width: 48px; height: 48px; flex-shrink: 0;
  background: #f0fdf4; border-radius: 50%;
  display: flex; align-items: center; justify-content: center;
}
.lcd-program-title { font-size: 0.95rem; font-weight: 700; margin: 0 0 5px; color: #111; }
.lcd-program-desc { font-size: 0.8rem; color: #666; line-height: 1.6; margin: 0; }

/* ─── CSR Target ─────────────────────────────────────────────── */
.section-target { padding: 60px 32px; background: #fff; }
.target-inner {
  max-width: 1000px; margin: 0 auto;
  display: grid; grid-template-columns: 1fr 360px;
  gap: 48px; align-items: center;
}
.target-title { font-size: 1.5rem; font-weight: 900; color: #111; margin: 0 0 8px; }
.target-divider { width: 40px; height: 3px; background: #16a34a; border-radius: 2px; margin-bottom: 10px; }
.target-subtitle { font-size: 0.83rem; color: #888; margin: 0 0 24px; }
.target-grid {
  display: grid; grid-template-columns: 1fr 1fr;
  gap: 16px 24px;
}
.target-item { display: flex; align-items: flex-start; gap: 12px; }
.target-icon {
  width: 40px; height: 40px; flex-shrink: 0;
  background: #f0fdf4; border-radius: 50%;
  display: flex; align-items: center; justify-content: center;
}
.target-item-text { font-size: 0.82rem; font-weight: 700; color: #111; margin: 0 0 2px; line-height: 1.4; }
.target-item-sub { font-size: 0.72rem; color: #888; margin: 0; line-height: 1.4; }
.target-right { }
.target-img {
  width: 100%; aspect-ratio: 4/5;
  object-fit: cover; border-radius: 16px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.12);
}

/* ─── Pilar / Tabs ───────────────────────────────────────────── */
.section-pilar { padding: 60px 32px; background: #f9fafb; }
.pilar-header { text-align: center; margin-bottom: 36px; }
.pilar-header h2 { font-size: 1.5rem; font-weight: 900; color: #111; margin: 0 0 10px; }
.pilar-header p { font-size: 0.85rem; color: #666; max-width: 460px; margin: 0 auto; line-height: 1.6; }

.tab-nav {
  display: flex; justify-content: center;
  border-bottom: 2px solid #e5e7eb;
  margin-bottom: 36px; gap: 0;
}
.tab-btn {
  display: flex; align-items: center; gap: 6px;
  padding: 12px 20px;
  font-size: 0.82rem; font-weight: 600; color: #999;
  background: none; border: none;
  border-bottom: 2px solid transparent; margin-bottom: -2px;
  cursor: pointer; white-space: nowrap;
  transition: color 0.2s, border-color 0.2s;
}
.tab-btn:hover { color: #16a34a; }
.tab-btn.active { color: #16a34a; border-bottom-color: #16a34a; }
.tab-icon { display: flex; align-items: center; }

.tab-content {
  max-width: 1000px; margin: 0 auto;
  display: grid; grid-template-columns: 1fr 1fr;
  gap: 48px; align-items: center;
}
.tab-left { }
.tab-image-wrap { border-radius: 16px; overflow: hidden; }
.tab-image {
  width: 100%; aspect-ratio: 4/3;
  object-fit: cover; display: block;
  transition: transform 0.4s ease;
}
.tab-image-wrap:hover .tab-image { transform: scale(1.03); }

/* ── Thumbnail gallery strip under the main tab image ───────── */
.tab-thumbs {
  display: flex; gap: 10px;
  margin-top: 14px;
  flex-wrap: wrap;
}
.tab-thumb-btn {
  width: 64px; height: 48px;
  border-radius: 8px; overflow: hidden;
  border: 2px solid transparent;
  padding: 0; cursor: pointer;
  background: none; flex-shrink: 0;
  opacity: 0.6;
  transition: opacity 0.2s ease, border-color 0.2s ease, transform 0.2s ease;
}
.tab-thumb-btn img {
  width: 100%; height: 100%;
  object-fit: cover; display: block;
}
.tab-thumb-btn:hover { opacity: 1; transform: translateY(-2px); }
.tab-thumb-btn--active {
  opacity: 1;
  border-color: #16a34a;
}

.tab-right { }
.tab-title { font-size: 1.4rem; font-weight: 900; color: #16a34a; margin: 0 0 8px; }
.tab-divider { width: 36px; height: 3px; background: #16a34a; border-radius: 2px; margin-bottom: 16px; }
.tab-desc { font-size: 0.85rem; color: #444; line-height: 1.75; margin: 0 0 20px; }
.tab-quote {
  display: flex; gap: 12px; align-items: flex-start;
  background: #f0fdf4; border-radius: 10px; padding: 16px;
}
.tab-quote p { font-size: 0.82rem; font-style: italic; color: #374151; line-height: 1.6; margin: 0; }

.tab-fade-enter-active, .tab-fade-leave-active { transition: opacity 0.22s ease, transform 0.22s ease; }
.tab-fade-enter-from { opacity: 0; transform: translateY(8px); }
.tab-fade-leave-to  { opacity: 0; transform: translateY(-8px); }

/* ─── CTA ────────────────────────────────────────────────────── */
.section-cta {
  position: relative; padding: 60px 32px;
  overflow: hidden;
}
.cta-bg {
  position: absolute; inset: 0;
  background: url('/images/csr-economy-2.jpg.jpeg') center/cover no-repeat;
  background-color: #111827;
  filter: brightness(0.25);
}
.cta-inner {
  position: relative; z-index: 2;
  display: flex; align-items: center;
  justify-content: space-between; gap: 32px;
}
.cta-text h2 { font-size: 1.4rem; font-weight: 800; color: #fff; margin: 0 0 8px; line-height: 1.4; }
.cta-text p { font-size: 0.82rem; color: rgba(255,255,255,0.65); margin: 0; line-height: 1.6; }
.cta-btn {
  flex-shrink: 0;
  display: flex; align-items: center; gap: 8px;
  background: transparent; color: #fff;
  border: 2px solid #fff;
  padding: 13px 26px; border-radius: 10px;
  font-size: 0.88rem; font-weight: 700;
  text-decoration: none; white-space: nowrap;
  transition: background 0.2s, color 0.2s, transform 0.2s;
}
.cta-btn:hover { background: #16a34a; border-color: #16a34a; transform: translateY(-2px); }

/* ─── Responsive ─────────────────────────────────────────────── */
@media (max-width: 768px) {
  .hero-content { padding: 0 24px 40px; }
  .hero-title { font-size: 2rem; }
  .lcd-card { grid-template-columns: 1fr; }
  .lcd-left { border-radius: 0; }
  .target-inner { grid-template-columns: 1fr; }
  .target-right { display: none; }
  .target-grid { grid-template-columns: 1fr; }
  .tab-content { grid-template-columns: 1fr; }
  .tab-nav { overflow-x: auto; justify-content: flex-start; }
  .cta-inner { flex-direction: column; text-align: center; }
}

/* ═══════════════════════════════════════════════════════════════
   ANIMATIONS
═══════════════════════════════════════════════════════════════ */

/* ── Scroll Reveal base ─────────────────────────────────────── */
.section-lcd,
.section-target,
.pilar-header,
.section-cta {
  opacity: 0;
  transform: translateY(36px);
  transition: opacity 0.7s cubic-bezier(0.22,1,0.36,1),
              transform 0.7s cubic-bezier(0.22,1,0.36,1);
}
.reveal-in {
  opacity: 1 !important;
  transform: translateY(0) !important;
}

/* ── LCD programs stagger ───────────────────────────────────── */
.lcd-program {
  opacity: 0;
  transform: translateX(-20px);
  transition: opacity 0.55s ease, transform 0.55s ease;
}
.lcd-program--visible {
  opacity: 1;
  transform: translateX(0);
}

/* ── LCD circle pulse ───────────────────────────────────────── */
.lcd-circle {
  animation: circlePulse 3s ease-in-out infinite;
}
@keyframes circlePulse {
  0%, 100% { box-shadow: 0 2px 12px rgba(22,163,74,0.12); }
  50%       { box-shadow: 0 4px 28px rgba(22,163,74,0.28); }
}

/* ── LCD leaf float ─────────────────────────────────────────── */
.lcd-leaf {
  animation: leafFloat 4s ease-in-out infinite;
}
@keyframes leafFloat {
  0%, 100% { transform: translateY(0) rotate(-5deg); }
  50%       { transform: translateY(-10px) rotate(5deg); }
}

/* ── CSR Target items stagger ───────────────────────────────── */
.target-item {
  opacity: 0;
  transform: translateY(16px);
  transition: opacity 0.5s ease, transform 0.5s ease;
}
.target-item--visible {
  opacity: 1;
  transform: translateY(0);
}

/* ── Target icon spin-in ─────────────────────────────────────── */
.target-icon {
  transition: transform 0.35s ease, background 0.3s ease;
}
.target-item--visible .target-icon {
  animation: iconPop 0.4s cubic-bezier(0.34,1.56,0.64,1) both;
}
@keyframes iconPop {
  from { transform: scale(0.5); opacity: 0; }
  to   { transform: scale(1); opacity: 1; }
}
.target-item:hover .target-icon {
  transform: scale(1.15) rotate(8deg);
  background: #dcfce7;
}

/* ── Target image reveal ─────────────────────────────────────── */
.target-img {
  opacity: 0;
  transform: scale(0.94) translateX(24px);
  transition: opacity 0.8s cubic-bezier(0.22,1,0.36,1) 0.3s,
              transform 0.8s cubic-bezier(0.22,1,0.36,1) 0.3s;
}
.img-reveal {
  opacity: 1 !important;
  transform: scale(1) translateX(0) !important;
}

/* ── Hero content word-by-word feel (stagger children) ───────── */
.hero-eyebrow { animation: fadeUp 0.6s 0.1s ease both; }
.hero-title   { animation: fadeUp 0.7s 0.25s ease both; }
.hero-desc    { animation: fadeUp 0.7s 0.4s ease both; }
.hero-btn     { animation: fadeUp 0.6s 0.55s ease both; }
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(22px); }
  to   { opacity: 1; transform: translateY(0); }
}

/* ── Tab underline slide ─────────────────────────────────────── */
.tab-btn { position: relative; overflow: hidden; }
.tab-btn::after {
  content: '';
  position: absolute; bottom: -2px; left: 50%; right: 50%;
  height: 2px; background: #16a34a;
  transition: left 0.3s ease, right 0.3s ease;
}
.tab-btn.active::after { left: 0; right: 0; }

/* ── LCD program hover ───────────────────────────────────────── */
.lcd-program {
  border-radius: 10px;
  transition: opacity 0.55s ease, transform 0.55s ease,
              background 0.2s ease, box-shadow 0.2s ease;
}
.lcd-program:hover {
  background: #f0fdf4;
  box-shadow: 0 2px 16px rgba(22,163,74,0.1);
}
.lcd-program-icon {
  transition: transform 0.3s ease, background 0.3s ease;
}
.lcd-program:hover .lcd-program-icon {
  transform: rotate(8deg) scale(1.1);
  background: #dcfce7;
}

/* ── CTA button shimmer ──────────────────────────────────────── */
/* ─── Footer Akhir ─────────────────────────────────────────── */
.final-footer {
  background-color: #4caf50; /* Warna hijau solid seperti contoh */
  padding: 40px 0;
  text-align: center;
  width: 100%;
  margin-top: 60px; /* Jarak dari konten di atasnya */
}
.footer-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 0 32px;
}
.final-footer p {
  color: #ffffff;
  font-size: 0.9rem;
  margin: 0;
  font-weight: 400;
}

/* ── Tab image zoom on change ────────────────────────────────── */
.tab-image-wrap {
  box-shadow: 0 8px 32px rgba(0,0,0,0.12);
}

/* ── Pilar header float-in ───────────────────────────────────── */
.pilar-header {
  transition-delay: 0.05s;
}

/* ── Section CTA bg parallax feel ───────────────────────────── */
.section-cta {
  transition-delay: 0.1s;
}
.section-cta.reveal-in .cta-bg {
  animation: bgScale 8s ease-out forwards;
}
@keyframes bgScale {
  from { transform: scale(1.06); }
  to   { transform: scale(1); }
}

/* ── Hero bg zoom override (already set, keep) ───────────────── */
.hero-bg { transform-origin: center center; }

</style>