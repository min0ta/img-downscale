const urls = [
    "https://m.media-amazon.com/images/I/71Qe7bz5bPL._UF894,1000_QL80_.jpg",
    "https://i.pinimg.com/736x/2d/95/e5/2d95e5886fc4c65a6778b5fee94a7d59.jpg",
    "https://img.freepik.com/free-photo/woman-beach-with-her-baby-enjoying-sunset_52683-144131.jpg",
    "https://img.freepik.com/free-photo/closeup-scarlet-macaw-from-side-view-scarlet-macaw-closeup-head_488145-3540.jpg",
    "https://burst.shopifycdn.com/photos/person-stands-on-rocks-poking-out-of-the-ocean-shoreline.jpg",
    "https://static.vecteezy.com/system/resources/thumbnails/009/273/280/small/concept-of-loneliness-and-disappointment-in-love-sad-man-sitting-element-of-the-picture-is-decorated-by-nasa-free-photo.jpg" 
]
const f = (url) => fetch("http://localhost:8080/tasks", {
    method: "POST",
    body: JSON.stringify({
        original_url: url
    })
})

urls.forEach(f)