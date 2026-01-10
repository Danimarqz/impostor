package game

import "impostor/internal/domain"

// DefaultCategories provides the static content for the game (English).
var DefaultCategories = []domain.Category{
	{
		Name: "General",
		Pairs: []domain.WordPair{
			{Real: "Hospital", Trap: "Pharmacy"},
			{Real: "Beach", Trap: "Pool"},
			{Real: "Cinema", Trap: "Theater"},
			{Real: "Library", Trap: "Bookstore"},
			{Real: "Airport", Trap: "Station"},
			{Real: "Guitar", Trap: "Violin"},
			{Real: "Coffee", Trap: "Tea"},
			{Real: "Sun", Trap: "Moon"},
            {Real: "Chair", Trap: "Stool"},
            {Real: "Laptop", Trap: "Tablet"},
            {Real: "Pen", Trap: "Pencil"},
            {Real: "Facebook", Trap: "Instagram"},
            {Real: "Google", Trap: "Bing"},
            {Real: "Marvel", Trap: "DC"},
            {Real: "Harry Potter", Trap: "Lord of the Rings"},
            {Real: "Star Wars", Trap: "Star Trek"},
            {Real: "Minecraft", Trap: "Roblox"},
            {Real: "Fortnite", Trap: "PUBG"},
            {Real: "Coca Cola", Trap: "Pepsi"},
            {Real: "McDonalds", Trap: "Burger King"},
            {Real: "Nike", Trap: "Adidas"},
            {Real: "iPhone", Trap: "Samsung"},
            {Real: "Windows", Trap: "MacOS"},
            {Real: "Python", Trap: "Java"},
            {Real: "Gold", Trap: "Silver"},
            {Real: "Diamond", Trap: "Ruby"},
            {Real: "Shirt", Trap: "T-Shirt"},
            {Real: "Shoes", Trap: "Sneakers"},
            {Real: "Glasses", Trap: "Sunglasses"},
            {Real: "Watch", Trap: "Bracelet"},
		},
	},
	{
		Name: "Animals",
		Pairs: []domain.WordPair{
			{Real: "Dog", Trap: "Wolf"},
			{Real: "Cat", Trap: "Tiger"},
			{Real: "Horse", Trap: "Zebra"},
			{Real: "Shark", Trap: "Dolphin"},
			{Real: "Eagle", Trap: "Falcon"},
			{Real: "Snake", Trap: "Lizard"},
            {Real: "Lion", Trap: "Cheetah"},
            {Real: "Bear", Trap: "Panda"},
            {Real: "Elephant", Trap: "Hippo"},
            {Real: "Giraffe", Trap: "Camel"},
            {Real: "Penguin", Trap: "Ostrich"},
            {Real: "Frog", Trap: "Toad"},
            {Real: "Bee", Trap: "Wasp"},
            {Real: "Ant", Trap: "Termite"},
            {Real: "Spider", Trap: "Scorpion"},
            {Real: "Butterfly", Trap: "Moth"},
            {Real: "Whale", Trap: "Orca"},
            {Real: "Crab", Trap: "Lobster"},
            {Real: "Octopus", Trap: "Squid"},
            {Real: "Rat", Trap: "Mouse"},
            {Real: "Rabbit", Trap: "Hare"},
            {Real: "Cow", Trap: "Bull"},
            {Real: "Sheep", Trap: "Goat"},
            {Real: "Chicken", Trap: "Turkey"},
            {Real: "Duck", Trap: "Goose"},
		},
	},
	{
		Name: "Food",
		Pairs: []domain.WordPair{
			{Real: "Pizza", Trap: "Burger"},
			{Real: "Sushi", Trap: "Sashimi"},
			{Real: "Tacos", Trap: "Burritos"},
			{Real: "Ice Cream", Trap: "Yogurt"},
			{Real: "Pasta", Trap: "Noodles"},
			{Real: "Cake", Trap: "Pie"},
            {Real: "Bread", Trap: "Toast"},
            {Real: "Butter", Trap: "Margarine"},
            {Real: "Cheese", Trap: "Cream"},
            {Real: "Milk", Trap: "Juice"},
            {Real: "Water", Trap: "Soda"},
            {Real: "Beer", Trap: "Wine"},
            {Real: "Whiskey", Trap: "Vodka"},
            {Real: "Tomato", Trap: "Potato"},
            {Real: "Onion", Trap: "Garlic"},
            {Real: "Apple", Trap: "Pear"},
            {Real: "Orange", Trap: "Lemon"},
            {Real: "Banana", Trap: "Plantain"},
            {Real: "Strawberry", Trap: "Raspberry"},
            {Real: "Grape", Trap: "Cherry"},
            {Real: "Chocolate", Trap: "Vanilla"},
            {Real: "Cookie", Trap: "Biscuit"},
            {Real: "Sandwich", Trap: "Wrap"},
            {Real: "Salad", Trap: "Soup"},
            {Real: "Steak", Trap: "Pork Chop"},
		},
	},
	{
		Name: "Places",
		Pairs: []domain.WordPair{
			{Real: "Paris", Trap: "Rome"},
			{Real: "New York", Trap: "Chicago"},
			{Real: "Tokyo", Trap: "Seoul"},
			{Real: "London", Trap: "Dublin"},
			{Real: "School", Trap: "University"},
			{Real: "Gym", Trap: "Park"},
            {Real: "Kitchen", Trap: "Bathroom"},
            {Real: "Bedroom", Trap: "Living Room"},
            {Real: "Hotel", Trap: "Motel"},
            {Real: "Restaurant", Trap: "Cafe"},
            {Real: "Bar", Trap: "Club"},
            {Real: "Library", Trap: "Museum"},
            {Real: "Zoo", Trap: "Aquarium"},
            {Real: "Beach", Trap: "Pool"},
            {Real: "Mountain", Trap: "Hill"},
            {Real: "Forest", Trap: "Jungle"},
            {Real: "Desert", Trap: "Canyon"},
            {Real: "Island", Trap: "Peninsula"},
            {Real: "Bridge", Trap: "Tunnel"},
            {Real: "Castle", Trap: "Palace"},
            {Real: "Pyramid", Trap: "Temple"},
            {Real: "Spain", Trap: "Italy"},
            {Real: "USA", Trap: "Canada"},
            {Real: "China", Trap: "Japan"},
            {Real: "Brazil", Trap: "Argentina"},
		},
	},
}

// DefaultCategoriesES provides Spanish translations
var DefaultCategoriesES = []domain.Category{
	{
		Name: "General",
		Pairs: []domain.WordPair{
			{Real: "Hospital", Trap: "Farmacia"},
			{Real: "Playa", Trap: "Piscina"},
			{Real: "Cine", Trap: "Teatro"},
			{Real: "Biblioteca", Trap: "Librería"},
			{Real: "Aeropuerto", Trap: "Estación"},
			{Real: "Guitarra", Trap: "Violín"},
			{Real: "Café", Trap: "Té"},
			{Real: "Sol", Trap: "Luna"},
            {Real: "Silla", Trap: "Taburete"},
            {Real: "Portátil", Trap: "Tableta"},
            {Real: "Bolígrafo", Trap: "Lápiz"},
            {Real: "Facebook", Trap: "Instagram"},
            {Real: "Google", Trap: "Bing"},
            {Real: "Marvel", Trap: "DC"},
            {Real: "Harry Potter", Trap: "El Señor de los Anillos"},
            {Real: "Star Wars", Trap: "Star Trek"},
            {Real: "Minecraft", Trap: "Roblox"},
            {Real: "Fortnite", Trap: "PUBG"},
            {Real: "Coca Cola", Trap: "Pepsi"},
            {Real: "McDonalds", Trap: "Burger King"},
            {Real: "Nike", Trap: "Adidas"},
            {Real: "iPhone", Trap: "Samsung"},
            {Real: "Windows", Trap: "MacOS"},
            {Real: "Python", Trap: "Java"},
            {Real: "Oro", Trap: "Plata"},
            {Real: "Diamante", Trap: "Rubí"},
            {Real: "Camisa", Trap: "Camiseta"},
            {Real: "Zapatos", Trap: "Zapatillas"},
            {Real: "Gafas", Trap: "Gafas de Sol"},
            {Real: "Reloj", Trap: "Pulsera"},
		},
	},
	{
		Name: "Animales",
		Pairs: []domain.WordPair{
			{Real: "Perro", Trap: "Lobo"},
			{Real: "Gato", Trap: "Tigre"},
			{Real: "Caballo", Trap: "Cebra"},
			{Real: "Tiburón", Trap: "Delfín"},
			{Real: "Águila", Trap: "Halcón"},
			{Real: "Serpiente", Trap: "Lagarto"},
            {Real: "León", Trap: "Guepardo"},
            {Real: "Oso", Trap: "Panda"},
            {Real: "Elefante", Trap: "Hipopótamo"},
            {Real: "Jirafa", Trap: "Camello"},
            {Real: "Pingüino", Trap: "Avestruz"},
            {Real: "Rana", Trap: "Sapo"},
            {Real: "Abeja", Trap: "Avispa"},
            {Real: "Hormiga", Trap: "Termita"},
            {Real: "Araña", Trap: "Escorpión"},
            {Real: "Mariposa", Trap: "Polilla"},
            {Real: "Ballena", Trap: "Orca"},
            {Real: "Cangrejo", Trap: "Langosta"},
            {Real: "Pulpo", Trap: "Calamar"},
            {Real: "Rata", Trap: "Ratón"},
            {Real: "Conejo", Trap: "Liebre"},
            {Real: "Vaca", Trap: "Toro"},
            {Real: "Oveja", Trap: "Cabra"},
            {Real: "Pollo", Trap: "Pavo"},
            {Real: "Pato", Trap: "Ganso"},
		},
	},
	{
		Name: "Comida",
		Pairs: []domain.WordPair{
			{Real: "Pizza", Trap: "Hamburguesa"},
			{Real: "Sushi", Trap: "Sashimi"},
			{Real: "Tacos", Trap: "Burritos"},
			{Real: "Helado", Trap: "Yogur"},
			{Real: "Pasta", Trap: "Fideos"},
			{Real: "Pastel", Trap: "Tarta"},
            {Real: "Pan", Trap: "Tostada"},
            {Real: "Mantequilla", Trap: "Margarina"},
            {Real: "Queso", Trap: "Crema"},
            {Real: "Leche", Trap: "Zumo"},
            {Real: "Agua", Trap: "Refresco"},
            {Real: "Cerveza", Trap: "Vino"},
            {Real: "Whisky", Trap: "Vodka"},
            {Real: "Tomate", Trap: "Patata"},
            {Real: "Cebolla", Trap: "Ajo"},
            {Real: "Manzana", Trap: "Pera"},
            {Real: "Naranja", Trap: "Limón"},
            {Real: "Plátano", Trap: "Plántano Macho"},
            {Real: "Fresa", Trap: "Frambuesa"},
            {Real: "Uva", Trap: "Cereza"},
            {Real: "Chocolate", Trap: "Vainilla"},
            {Real: "Galleta", Trap: "Bizcocho"},
            {Real: "Sándwich", Trap: "Wrap"},
            {Real: "Ensalada", Trap: "Sopa"},
            {Real: "Filete", Trap: "Chuleta de Cerdo"},
		},
	},
	{
		Name: "Lugares",
		Pairs: []domain.WordPair{
			{Real: "París", Trap: "Roma"},
			{Real: "Nueva York", Trap: "Chicago"},
			{Real: "Tokio", Trap: "Seúl"},
			{Real: "Londres", Trap: "Dublín"},
			{Real: "Escuela", Trap: "Universidad"},
			{Real: "Gimnasio", Trap: "Parque"},
            {Real: "Cocina", Trap: "Baño"},
            {Real: "Dormitorio", Trap: "Salón"},
            {Real: "Hotel", Trap: "Motel"},
            {Real: "Restaurante", Trap: "Cafetería"},
            {Real: "Bar", Trap: "Club"},
            {Real: "Biblioteca", Trap: "Museo"},
            {Real: "Zoológico", Trap: "Acuario"},
            {Real: "Playa", Trap: "Piscina"},
            {Real: "Montaña", Trap: "Colina"},
            {Real: "Bosque", Trap: "Selva"},
            {Real: "Desierto", Trap: "Cañón"},
            {Real: "Isla", Trap: "Península"},
            {Real: "Puente", Trap: "Túnel"},
            {Real: "Castillo", Trap: "Palacio"},
            {Real: "Pirámide", Trap: "Templo"},
            {Real: "España", Trap: "Italia"},
            {Real: "EE.UU.", Trap: "Canadá"},
            {Real: "China", Trap: "Japón"},
            {Real: "Brasil", Trap: "Argentina"},
		},
	},
}

func GetCategoryByName(name string, language string) domain.Category {
    if name == "✨ Infinite" {
        return domain.Category{Name: "✨ Infinite", Pairs: []domain.WordPair{}}
    }

	categories := DefaultCategories
	if language == "es" {
		categories = DefaultCategoriesES
	}

	for _, c := range categories {
		if c.Name == name {
			return c
		}
	}
	// Fallback to General if not found
	return categories[0]
}

func GetAllCategoryNames(language string) []string {
	categories := DefaultCategories
	if language == "es" {
		categories = DefaultCategoriesES
	}

	names := make([]string, 0, len(categories)+1)
    names = append(names, "✨ Infinite")
	for _, c := range categories {
		names = append(names, c.Name)
	}
	return names
}
