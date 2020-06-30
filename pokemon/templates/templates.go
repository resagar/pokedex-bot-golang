package templates

const TemplateListPokemon string = `
Conteo: {{.Count}}
Lista de Pokemons:
{{range .Results}}/{{.Name}},{{end}}
`
const TemplateOnePokemon string = `
Imagen: {{.Sprite.FrontDefault}}
Orden: {{.Order}}
Name: {{.Name}}
Especie: {{.Species.Name}}
Tipo: 
{{range .Types}}{{.Type.Name}}, {{end}}
Habilidades: 
{{range .Abilities}}{{.Ability.Name}}, {{end}}
Movimientos: 
{{range .Moves}}{{.Move.Name}}, {{end}}
Versiones donde Aparece:
{{range .GameIndex}}{{.Version.Name}}, {{end}}
`