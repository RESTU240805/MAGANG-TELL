<template>

  <section class="relative h-[480px] flex items-center justify-center bg-cover bg-center overflow-hidden"
    style="background-image: url('/images/pulp.jpeg')">

    <div class="absolute inset-0 bg-gradient-to-b from-black/60 via-black/50 to-black/80"></div>
    <div class="absolute inset-0 hero-grid opacity-10"></div>

    <div class="orb orb-1"></div>
    <div class="orb orb-2"></div>
    <div class="orb orb-3"></div>

    <div class="relative z-10 text-center px-6 w-full max-w-4xl mx-auto">
      <h1 class="text-5xl md:text-6xl font-black text-white leading-tight mb-2 min-h-[72px]">
        <span class="typewriter-text">{{ displayedTitle }}</span><span
          class="typewriter-cursor" :class="{ 'cursor-hidden': typingDone }">|</span>
      </h1>

      <div class="flex items-center justify-center mt-7 mb-5 opacity-0 anim-fade-up" style="animation-delay:2.6s">
        <div class="divider-line"></div>
      </div>

      <div class="flex items-center justify-center gap-2 text-xs text-gray-400 opacity-0 anim-fade-up" style="animation-delay:2.8s">
        <RouterLink to="/" class="hover:text-white transition">Home</RouterLink>
        <span class="text-green-500">/</span>
        <span class="hover:text-white transition cursor-pointer">Sustainability</span>
        <span class="text-green-500">/</span>
        <span class="text-green-400 font-semibold">Supply Chain Management</span>
      </div>

      <div class="absolute bottom-8 left-1/2 -translate-x-1/2 opacity-0 anim-fade-up" style="animation-delay:3.1s">
        <div class="scroll-indicator">
          <div class="scroll-dot"></div>
        </div>
        <p class="text-gray-500 text-[10px] tracking-widest uppercase mt-2">Scroll</p>
      </div>
    </div>
  </section>

  <section class="py-20 bg-gray-50">
    <div class="max-w-5xl mx-auto px-10">

      <div class="text-center mb-14 reveal" data-delay="0">
        <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-4 py-1.5 rounded-full mb-4">
          
        </span>
        <h2 class="text-3xl font-black text-gray-900">Supply Chain Pillars</h2>
        <p class="text-gray-400 text-sm mt-2 max-w-md mx-auto">
          Click each section to explore our complete supply chain management framework.
        </p>
        <div class="flex items-center justify-center gap-2 mt-5">
          <div class="w-8 h-0.5 bg-green-300 rounded-full"></div>
          <div class="w-2 h-2 bg-green-500 rounded-full"></div>
          <div class="w-8 h-0.5 bg-green-300 rounded-full"></div>
        </div>
      </div>

      <div class="space-y-4">

        <div class="bg-white rounded-2xl border border-gray-100 shadow-sm overflow-hidden reveal" data-delay="100">
          <button
            @click="toggle(0)"
            class="w-full flex items-center justify-between px-7 py-5 text-left transition-all duration-300"
            :class="open[0] ? 'bg-gradient-to-r from-green-600 to-green-500' : 'bg-white hover:bg-gray-50'">
            <div class="flex items-center gap-4">
              <span class="w-8 h-8 rounded-xl flex items-center justify-center text-sm font-black transition-all"
                :class="open[0] ? 'bg-white/20 text-white' : 'bg-green-100 text-green-700'">①</span>
              <span :class="open[0] ? 'text-white' : 'text-gray-900'" class="font-bold text-base transition">
                Strategy and Strong Management Plan
              </span>
            </div>
            <svg :class="[open[0] ? 'rotate-180 text-white' : 'text-gray-400', 'transition-transform duration-300 w-5 h-5 flex-shrink-0']"
              fill="none" stroke="currentColor" stroke-width="2.5" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>
          <transition name="accordion">
            <div v-if="open[0]" class="px-7 pb-7 pt-4">
              <p class="text-gray-500 text-sm leading-relaxed mb-5">We have a strategy for strong management plan as follows:</p>
              <div class="bg-green-50 border border-green-100 rounded-xl p-5 mb-4">
                <div class="flex items-start gap-3 mb-3">
                  <span class="w-2 h-2 rounded-full bg-green-500 mt-1.5 flex-shrink-0"></span>
                  <p class="text-gray-700 text-sm font-semibold">To keep stable pulp production and its good quality which consist of:</p>
                </div>
                <div class="grid grid-cols-2 gap-2 ml-5">
                  <div v-for="(item, idx) in strategyItems" :key="item"
                    class="flex items-center gap-2 bg-white rounded-lg px-4 py-2.5 border border-green-100 shadow-sm stagger-item"
                    :style="`animation-delay:${idx * 0.1}s`">
                    <span class="w-1.5 h-1.5 rounded-full bg-green-400 flex-shrink-0"></span>
                    <span class="text-sm text-gray-600">{{ item }}</span>
                  </div>
                </div>
              </div>
              <div class="space-y-3">
                <div class="flex items-start gap-3 bg-gray-50 rounded-xl px-5 py-3.5 border border-gray-100 stagger-item" style="animation-delay:0.4s">
                  <span class="w-2 h-2 rounded-full bg-green-500 mt-1.5 flex-shrink-0"></span>
                  <p class="text-gray-600 text-sm">To implement sound Industrial Relation, Corporate Social Responsibility.</p>
                </div>
                <div class="flex items-start gap-3 bg-gray-50 rounded-xl px-5 py-3.5 border border-gray-100 stagger-item" style="animation-delay:0.5s">
                  <span class="w-2 h-2 rounded-full bg-green-500 mt-1.5 flex-shrink-0"></span>
                  <p class="text-gray-600 text-sm">To develop competent and motivated People.</p>
                </div>
              </div>
            </div>
          </transition>
        </div>

        <div class="bg-white rounded-2xl border border-gray-100 shadow-sm overflow-hidden reveal" data-delay="200">
          <button
            @click="toggle(1)"
            class="w-full flex items-center justify-between px-7 py-5 text-left transition-all duration-300"
            :class="open[1] ? 'bg-gradient-to-r from-green-600 to-green-500' : 'bg-white hover:bg-gray-50'">
            <div class="flex items-center gap-4">
              <span class="w-8 h-8 rounded-xl flex items-center justify-center text-sm font-black transition-all"
                :class="open[1] ? 'bg-white/20 text-white' : 'bg-green-100 text-green-700'">②</span>
              <span :class="open[1] ? 'text-white' : 'text-gray-900'" class="font-bold text-base transition">
                Supply Chain Sustainability
              </span>
            </div>
            <svg :class="[open[1] ? 'rotate-180 text-white' : 'text-gray-400', 'transition-transform duration-300 w-5 h-5 flex-shrink-0']"
              fill="none" stroke="currentColor" stroke-width="2.5" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>
          <transition name="accordion">
            <div v-if="open[1]" class="px-7 pb-7 pt-4">
              <div class="bg-gradient-to-r from-green-800 to-green-700 rounded-xl p-5 mb-5 text-center stagger-item">
                <p class="text-green-300 text-xs font-bold tracking-widest uppercase mb-1">TEL Motto</p>
                <p class="text-white text-lg font-black">"Fairness, Innovation & Harmony"</p>
              </div>
              <div class="space-y-4">
                <div v-for="(item, i) in sustainabilityItems" :key="i"
                  class="flex gap-4 bg-gray-50 rounded-xl p-5 border border-gray-100 hover:border-green-200 hover:bg-green-50/50 transition stagger-item"
                  :style="`animation-delay:${i * 0.1}s`">
                  <div class="w-8 h-8 rounded-full bg-green-600 text-white font-black text-xs flex items-center justify-center flex-shrink-0 mt-0.5">
                    {{ i + 1 }}
                  </div>
                  <p class="text-gray-600 text-sm leading-relaxed">{{ item }}</p>
                </div>
              </div>
            </div>
          </transition>
        </div>

        <div class="bg-white rounded-2xl border border-gray-100 shadow-sm overflow-hidden reveal" data-delay="300">
          <button
            @click="toggle(2)"
            class="w-full flex items-center justify-between px-7 py-5 text-left transition-all duration-300"
            :class="open[2] ? 'bg-gradient-to-r from-green-600 to-green-500' : 'bg-white hover:bg-gray-50'">
            <div class="flex items-center gap-4">
              <span class="w-8 h-8 rounded-xl flex items-center justify-center text-sm font-black transition-all"
                :class="open[2] ? 'bg-white/20 text-white' : 'bg-green-100 text-green-700'">③</span>
              <span :class="open[2] ? 'text-white' : 'text-gray-900'" class="font-bold text-base transition">
                Policy & Guidelines
              </span>
            </div>
            <svg :class="[open[2] ? 'rotate-180 text-white' : 'text-gray-400', 'transition-transform duration-300 w-5 h-5 flex-shrink-0']"
              fill="none" stroke="currentColor" stroke-width="2.5" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>
          <transition name="accordion">
            <div v-if="open[2]" class="px-7 pb-7 pt-4">
              <div class="space-y-5">
                <div v-for="(policy, i) in policies" :key="i"
                  class="border border-gray-100 rounded-xl overflow-hidden stagger-item"
                  :style="`animation-delay:${i * 0.07}s`">
                  <div class="flex items-center gap-3 bg-gray-50 px-5 py-3 border-b border-gray-100">
                    <span class="w-7 h-7 rounded-lg bg-green-600 text-white font-black text-xs flex items-center justify-center flex-shrink-0">{{ i + 1 }}</span>
                    <h4 class="font-bold text-gray-900 text-sm">{{ policy.title }}</h4>
                  </div>
                  <div class="px-5 py-4 space-y-2">
                    <div v-for="(point, j) in policy.points" :key="j" class="flex items-start gap-2.5">
                      <span class="w-1.5 h-1.5 rounded-full bg-green-400 mt-1.5 flex-shrink-0"></span>
                      <p class="text-gray-500 text-sm leading-relaxed">{{ point }}</p>
                    </div>
                    <div v-if="policy.procedures" class="mt-4 ml-4 border-l-2 border-green-200 pl-4 space-y-3">
                      <p class="text-gray-600 text-xs font-semibold uppercase tracking-widest mb-2">Procedures for vendors not meeting labor standards:</p>
                      <div v-for="(proc, k) in policy.procedures" :key="k" class="flex gap-3 bg-green-50 rounded-lg p-3">
                        <span class="w-5 h-5 rounded-full bg-green-600 text-white font-bold text-[10px] flex items-center justify-center flex-shrink-0 mt-0.5">{{ k + 1 }}</span>
                        <p class="text-gray-600 text-xs leading-relaxed">{{ proc }}</p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </transition>
        </div>

      </div>
    </div>
  </section>

  <footer class="bg-green-600 text-white text-[13px] py-4">
    <div class="max-w-5xl mx-auto px-10 flex items-center justify-between">
      <p class="font-medium tracking-wide">
        Copyright 2019 PT TELPP. All right reserved.
      </p>
    </div>
  </footer>

</template>

<script setup>
import { ref, onMounted } from 'vue'

// ── Accordion ──
const open = ref([true, false, false])
function toggle(i) { open.value[i] = !open.value[i] }

// ── Typewriter ──
const fullTitle = 'Supply Chain Management'
const displayedTitle = ref('')
const typingDone = ref(false)

onMounted(() => {
  // Typewriter
  let i = 0
  setTimeout(() => {
    const iv = setInterval(() => {
      if (i < fullTitle.length) {
        displayedTitle.value += fullTitle[i++]
      } else {
        clearInterval(iv)
        setTimeout(() => { typingDone.value = true }, 900)
      }
    }, 58)
  }, 500)

  // Intersection Observer — animasi saat scroll ke bawah
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        const delay = parseInt(entry.target.dataset.delay || 0)
        setTimeout(() => {
          entry.target.classList.add('revealed')
        }, delay)
        observer.unobserve(entry.target)
      }
    })
  }, { threshold: 0.15 })

  // Observe semua elemen .reveal dan .reveal-cta
  document.querySelectorAll('.reveal, .reveal-cta').forEach(el => observer.observe(el))
})

// ── Data ──
const strategyItems = [
  'Log Management',
  'Procurement Management',
  'Spare Part Management',
  'Contract Management',
]

const sustainabilityItems = [
  'PT.Tanjungenim Lestari Pulp and Paper has a principle stated in a motto "Fairness, Innovation & Harmony" which underlines all of its activities and transactions that occur to ensure the sustainability of business activities in harmony.',
  'PT.Tanjungenim Lestari Pulp and Paper commits to uphold Good Corporate Governance (GCG) principles in its operation. To achieve this, the participation of all stakeholders including suppliers is very essential.',
  'PT.Tanjungenim Lestari Pulp and Paper is part of Marubeni Group who already set its supply chain sustainability guidelines, supporting an environmentally friendly, healthy and sustainable society.',
  'Cascading the Marubeni Group supply chain sustainability guidelines, PT.Tanjungenim Lestari Pulp and Paper asks for the understanding and cooperation of its business partners in observing the Guidelines.',
]

const policies = [
  { title: 'Observance of Laws', points: ['Observe the laws of the countries where business is conducted and laws relating to business transactions.'] },
  {
    title: 'Respect for Human Rights',
    points: [
      'Respect human rights without discrimination, harassment of any kind, abuse or other inhumane treatment.',
      'No child labor or forced labor.',
      'Proper management of employees\' work hours, breaks and holidays.',
      'Payment of the legally mandated minimum wage.',
      'Respect for employees\' right to unionize.',
    ],
  },
  { title: 'Conservation of the Environment', points: ['Recognize that climate change issues are important and respond appropriately.', 'Protect the natural environment.', 'Reduce environmental negative impact; prevent pollution.'] },
  { title: 'Fair Transactions', points: ['Conduct fair transactions and do not inhibit free competition.', 'Prevent corruption; offer no bribes or illegal contributions.'] },
  { title: 'Safety and Health', points: ['Ensure safe and healthy workplaces and maintain a good working environment.'] },
  { title: 'Quality Control', points: ['Maintain the quality and safety of products and services.'] },
  {
    title: 'Information Disclosure',
    points: ['Timely and appropriate disclosure of information.', 'The Marubeni Group has set out procedures for dealing with vendors that do not meet labor standards:'],
    procedures: [
      'When it comes to light that a vendor has failed to meet labor standards, we will ask the vendor to ascertain the facts and prepare a report on the background of the issue and improvement measures.',
      'If we determine that improvement measures are insufficient, we will request that further measures be taken.',
      'If, despite implementing steps (i) and (ii) above, the situation does not improve, we will examine whether to continue our relationship with the vendor.',
    ],
  },
]
</script>

<style scoped>

/* ══════════════════════════════════════
   SCROLL REVEAL — state awal & revealed
   ══════════════════════════════════════ */
.reveal {
  opacity: 0;
  transform: translateY(36px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1),
              transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.reveal.revealed {
  opacity: 1;
  transform: translateY(0);
}

/* CTA: muncul dari bawah + sedikit scale */
.reveal-cta {
  opacity: 0;
  transform: translateY(28px) scale(0.97);
  transition: opacity 0.65s cubic-bezier(0.22, 1, 0.36, 1),
              transform 0.65s cubic-bezier(0.22, 1, 0.36, 1);
}
.reveal-cta.revealed {
  opacity: 1;
  transform: translateY(0) scale(1);
}

/* ══════════════════════════════
   CTA section extras
   ══════════════════════════════ */
.cta-ring {
  position: absolute;
  border-radius: 50%;
  border: 1px solid rgba(74, 222, 128, 0.12);
  pointer-events: none;
  animation: ringPulse 6s ease-in-out infinite;
}
.cta-ring-1 { width: 500px; height: 500px; top: 50%; left: 50%; transform: translate(-50%, -50%); }
.cta-ring-2 { width: 700px; height: 700px; top: 50%; left: 50%; transform: translate(-50%, -50%); animation-delay: 1.5s; }
@keyframes ringPulse {
  0%, 100% { opacity: 0.4; transform: translate(-50%, -50%) scale(1); }
  50%       { opacity: 0.8; transform: translate(-50%, -50%) scale(1.04); }
}

.cta-pill {
  display: flex;
  align-items: center;
  gap: 12px;
  background: rgba(255, 255, 255, 0.08);
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: 14px;
  padding: 14px 20px;
  transition: background 0.25s, transform 0.25s, border-color 0.25s;
  cursor: default;
  min-width: 180px;
}
.cta-pill:hover {
  background: rgba(255, 255, 255, 0.15);
  border-color: rgba(74, 222, 128, 0.4);
  transform: translateY(-3px);
}

.pulse-dot-sm {
  animation: dotPulse 1.8s ease-in-out infinite;
}
@keyframes dotPulse {
  0%, 100% { opacity: 0.6; transform: scale(1); }
  50%       { opacity: 1;   transform: scale(1.5); }
}

/* ══════════════════════════════
   HERO
   ══════════════════════════════ */
.hero-grid {
  background-image:
    linear-gradient(rgba(74,222,128,0.3) 1px, transparent 1px),
    linear-gradient(90deg, rgba(74,222,128,0.3) 1px, transparent 1px);
  background-size: 60px 60px;
  animation: gridMove 20s linear infinite;
}
@keyframes gridMove {
  from { background-position: 0 0; }
  to   { background-position: 60px 60px; }
}

.orb { position:absolute; border-radius:50%; filter:blur(60px); pointer-events:none; }
.orb-1 { width:300px; height:300px; background:rgba(74,222,128,0.12); top:-80px; left:-80px; animation:orbFloat 8s ease-in-out infinite; }
.orb-2 { width:200px; height:200px; background:rgba(21,128,61,0.2); bottom:0; right:10%; animation:orbFloat 10s ease-in-out infinite reverse; }
.orb-3 { width:150px; height:150px; background:rgba(74,222,128,0.08); top:30%; left:50%; animation:orbFloat 12s ease-in-out infinite; }
@keyframes orbFloat {
  0%,100% { transform:translate(0,0) scale(1); }
  33%     { transform:translate(20px,-20px) scale(1.05); }
  66%     { transform:translate(-15px,15px) scale(0.97); }
}

.divider-line {
  height:1px; width:0;
  background:linear-gradient(to right, transparent, #4ade80, transparent);
  animation:lineGrow 1s 2.6s cubic-bezier(0.22,1,0.36,1) forwards;
}
@keyframes lineGrow {
  from { width:0; opacity:0; }
  to   { width:120px; opacity:1; }
}

.scroll-indicator {
  width:22px; height:36px;
  border:2px solid rgba(255,255,255,0.3);
  border-radius:11px; margin:0 auto;
  display:flex; align-items:flex-start; justify-content:center; padding-top:5px;
}
.scroll-dot {
  width:4px; height:8px; background:#4ade80; border-radius:2px;
  animation:scrollBounce 1.8s ease-in-out infinite;
}
@keyframes scrollBounce {
  0%,100% { transform:translateY(0); opacity:1; }
  50%     { transform:translateY(10px); opacity:0.4; }
}

.typewriter-cursor {
  display:inline-block; color:#4ade80; font-weight:300; margin-left:3px;
  animation:blink 0.65s step-end infinite;
}
.typewriter-cursor.cursor-hidden { opacity:0; animation:none; }
@keyframes blink {
  0%,100% { opacity:1; }
  50%     { opacity:0; }
}

/* ══════════════════════════════
   GENERAL ENTRANCE
   ══════════════════════════════ */
@keyframes fadeUp {
  from { opacity:0; transform:translateY(28px); }
  to   { opacity:1; transform:translateY(0); }
}
.anim-fade-up { animation:fadeUp 0.8s cubic-bezier(0.22,1,0.36,1) forwards; }

.stagger-item {
  opacity:0;
  animation:fadeUp 0.5s cubic-bezier(0.22,1,0.36,1) forwards;
}

/* ── Accordion transition ── */
.accordion-enter-active,
.accordion-leave-active { transition:all 0.38s cubic-bezier(0.22,1,0.36,1); overflow:hidden; }
.accordion-enter-from,
.accordion-leave-to  { opacity:0; max-height:0; }
.accordion-enter-to,
.accordion-leave-from { opacity:1; max-height:2000px; }
</style>