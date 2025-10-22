document.getElementById('search-form').addEventListener('submit', async function(e) {
    e.preventDefault(); // impede o reload da página

    const username = document.getElementById('username').value.trim();
    const infoDiv = document.getElementById('user-info');
    const notFound = document.getElementById('not-found');

    // esconde mensagens antigas
    infoDiv.style.display = 'none';
    notFound.style.display = 'none';

    if (!username) return;

    try {
        // faz requisição para o seu backend Go
        const response = await fetch(`/api/user?username=${username}`);

        if (!response.ok) {
            notFound.style.display = 'block';
            return;
        }

        const user = await response.json();

        // preenche os dados retornados
        document.getElementById('avatar').src = user.avatar_url;
        document.getElementById('name').textContent = user.name || '—';
        document.getElementById('login').textContent = user.login || '—';
        document.getElementById('bio').textContent = user.bio || '—';
        document.getElementById('repos').textContent = user.public_repos || 0;
        document.getElementById('followers').textContent = user.followers || 0;
        document.getElementById('following').textContent = user.following || 0;

        infoDiv.style.display = 'block';
    } catch (err) {
        notFound.style.display = 'block';
        console.error(err);
    }
});