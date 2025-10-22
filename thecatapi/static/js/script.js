const btn = document.getElementById('gerarBtn');
const img = document.getElementById('catImage');
const loading = document.getElementById('loading');

btn.addEventListener('click', async () => {
    img.style.display = 'none';
    loading.style.display = 'block';
    loading.textContent = 'Carregando...';

    try {
        const response = await fetch('/api/cat');
        if (!response.ok) throw new Error('Erro ao buscar gato');

        const data = await response.json();
        img.src = data.image;
        img.onload = () => {
            loading.style.display = 'none';
            img.style.display = 'block';
        };
    } catch (err) {
        loading.textContent = 'Falha ao carregar imagem 😿';
        console.error(err);
    }
});