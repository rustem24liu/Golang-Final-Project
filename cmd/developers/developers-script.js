document.addEventListener("DOMContentLoaded", function() {
    const developers = [
        { name: "Suleimenov Ayan", role: "Team Leader" },
        { name: "Bagytzhan Zhalgas" },
        { name: "Tolegen Nursultan" },
        { name: "Rustem" },
        { name: "Liu Rustem" }
    ];

    const developerContainer = document.querySelector('.developers-container');
    const developerDetailsContainer = document.querySelector('.developer-details-container');

    // Populate developers list
    developers.forEach((developer, index) => {
        const developerElement = document.createElement('div');
        developerElement.classList.add('developer');
        developerElement.textContent = developer.name;
        developerElement.addEventListener('click', () => showDeveloperDetails(developer, index));
        developerContainer.appendChild(developerElement);
    });

    // Function to show developer details
    function showDeveloperDetails(developer, index) {
        const developerNameElement = document.getElementById('developer-name');
        const developerImagesContainer = document.querySelector('.developer-images-container');
        const paginationContainer = document.querySelector('.pagination-container');

        developerNameElement.textContent = `${developer.name} - ${developer.role || 'Developer'}`;
        developerImagesContainer.innerHTML = ''; // Clear previous images

        const images = [
            "https://unsplash.com/photos/VgQXtsaivtc",
            "https://unsplash.com/photos/jHrvMzBVI-U",
            "https://unsplash.com/photos/1ZCk_LNMcDM",
            "https://unsplash.com/photos/12ojQ9eFfRc",
            "https://unsplash.com/photos/EvXTZ7fFL4A",
            "https://unsplash.com/photos/4CLeGhS61K0",
            "https://unsplash.com/photos/ruRoiEw_Z1A",
            "https://unsplash.com/photos/x9mXsWdYbuc",
            "https://unsplash.com/photos/lvD-_XsJFkY",
            "https://unsplash.com/photos/XevE8mU-nKo",
            "https://unsplash.com/photos/C_KGY9uNIxk",
            "https://unsplash.com/photos/Lb93Y1jT9rg"
        ];

        // Calculate pagination indices
        const pageSize = 3;
        const startIndex = index * pageSize;
        const endIndex = Math.min(startIndex + pageSize, images.length);

        // Display images for the current developer with pagination
        for (let i = startIndex; i < endIndex; i++) {
            const imageElement = document.createElement('img');
            imageElement.src = images[i];
            developerImagesContainer.appendChild(imageElement);
        }

        // Pagination
        paginationContainer.innerHTML = '';
        if (startIndex > 0) {
            const prevButton = createPaginationButton('Prev');
            prevButton.addEventListener('click', () => showDeveloperDetails(developer, index - 1));
            paginationContainer.appendChild(prevButton);
        }
        if (endIndex < images.length) {
            const nextButton = createPaginationButton('Next');
            nextButton.addEventListener('click', () => showDeveloperDetails(developer, index + 1));
            paginationContainer.appendChild(nextButton);
        }

        // Show developer details container
        developerDetailsContainer.style.display = 'block';
    }

    // Function to create pagination button
    function createPaginationButton(text) {
        const button = document.createElement('button');
        button.textContent = text;
        return button;
    }
});
