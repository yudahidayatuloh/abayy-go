package utils

func GetFooter() string {
	return `
	<footer class="bg-white mt-20 border-t-1 border-gray-100 shadow-t-lg">
		<div class="max-w-7xl mx-auto px-6 py-8 grid grid-cols-1 md:grid-cols-3 gap-8">
			
			<!-- Brand -->
			<div>
			<h2 class="text-xl font-bold">SMK Negeri 4 Tasikmalaya</h2>
			<p class="mt-2 text-sm">
				Menghasilkan lulusan yang Cerdas, Akhlak mulia, Kreatif, Aktif, Produktif, 
				dengan berlandaskan Iman dan Taqwa (CAKAP BERIMTAQ).
			</p>
			</div>

			<!-- Navigasi -->
			<div class="flex flex-col space-y-2">
			<h3 class="text-lg font-semibold">Navigasi</h3>
			<a href="/" class="hover:text-[#0265af]">Beranda</a>
			<a href="/profil" class="hover:text-[#0265af]">Profil</a>
			<a href="/eskul" class="hover:text-[#0265af]">Ekstrakurikuler</a>
			<a href="/galeri" class="hover:text-[#0265af]">Galeri</a>
			</div>

			<!-- Kontak -->
			<div>
			<h3 class="text-lg font-semibold">Alamat</h3>
			<p class="text-sm mt-2">Jl. Depok, Sukamenak, Kec. Purbaratu, Kab. Tasikmalaya, Jawa Barat 46196</p>
			</div>
		</div>

		<!-- Copyright -->
		<div class="border-t border-gray-200 text-center py-4 text-sm text-gray-600">
			© 2025 SMK Negeri 4 Tasikmalaya. All rights reserved.
		</div>
	</footer>

	`
}