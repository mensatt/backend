package seeds

import (
	"context"
	"github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/database/schema"
)

func seedTags(ctx context.Context, client *ent.Client) error {
	err := client.Tag.CreateBulk(
		client.Tag.Create().SetID("1").SetName("Farbstoff").SetDescription("mit Farbstoff"),
		client.Tag.Create().SetID("2").SetName("Koffein").SetDescription("mit Koffein"),
		client.Tag.Create().SetID("4").SetName("Konservierungsstoffen").SetDescription("mit Konservierungsstoffen"),
		client.Tag.Create().SetID("5").SetName("Süßungsmittel").SetDescription("mit Süßungsmittel"),
		client.Tag.Create().SetID("7").SetName("Antioxidationsmittel").SetDescription("mit Antioxidationsmittel"),
		client.Tag.Create().SetID("8").SetName("Geschmacksverstärker").SetDescription("mit Geschmacksverstärker"),
		client.Tag.Create().SetID("9").SetName("geschwefelt").SetDescription("geschwefelt"),
		client.Tag.Create().SetID("10").SetName("geschwärzt").SetDescription("geschwärzt"),
		client.Tag.Create().SetID("11").SetName("gewachst").SetDescription("gewachst"),
		client.Tag.Create().SetID("12").SetName("Phosphat").SetDescription("mit Phosphat"),
		client.Tag.Create().SetID("13").SetName("Phenylalaninquelle").SetDescription("mit einer Phenylalaninquelle"),
		client.Tag.Create().SetID("30").SetName("Fettglasur").SetDescription("mit Fettglasur"),

		client.Tag.Create().SetID("S").SetName("Schwein").SetDescription("Schweinefleisch").SetShortName("🐷").SetPriority(schema.High),
		client.Tag.Create().SetID("R").SetName("Rind").SetDescription("Rindfleisch").SetShortName("🐮").SetPriority(schema.High),
		client.Tag.Create().SetID("G").SetName("Geflügel").SetDescription("Geflügel").SetShortName("🐔").SetPriority(schema.High),
		client.Tag.Create().SetID("L").SetName("Lamm").SetDescription("Lamm").SetShortName("🐑").SetPriority(schema.High),
		client.Tag.Create().SetID("W").SetName("Wild").SetDescription("Wildfleisch").SetShortName("🦌").SetPriority(schema.High),
		client.Tag.Create().SetID("F").SetName("Fisch").SetDescription("Fisch").SetShortName("🐟").SetPriority(schema.High),
		client.Tag.Create().SetID("V").SetName("Vegetarisch").SetDescription("Vegetarisch").SetShortName("🥕").SetPriority(schema.High),
		client.Tag.Create().SetID("Veg").SetName("Vegan").SetDescription("Vegan").SetShortName("🌱").SetPriority(schema.High),
		client.Tag.Create().SetID("MSC").SetName("MSC Fisch").SetDescription("MSC Fisch").SetShortName("🐟").SetPriority(schema.High),
		client.Tag.Create().SetID("Gf").SetName("Glutenfrei").SetDescription("Glutenfrei").SetShortName("Glutenfrei").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetID("CO2").SetName("CO2-Neutral").SetDescription("CO2-Neutral").SetShortName("CO2-Neutral").SetPriority(schema.Medium),
		client.Tag.Create().SetID("Bio").SetName("Bio").SetDescription("Bio").SetShortName("Bio").SetPriority(schema.Medium),
		client.Tag.Create().SetID("MV").SetName("MensaVital").SetDescription("MensaVital").SetShortName("MensaVital").SetPriority(schema.Medium),
		// todo: missing alchohol icon

		client.Tag.Create().SetID("Wz").SetName("Weizen").SetDescription("Weizen (Gluten)").SetShortName("🌾").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetID("Ro").SetName("Roggen").SetDescription("Roggen (Gluten)").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetID("Ge").SetName("Gerste").SetDescription("Gerste (Gluten)").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetID("Hf").SetName("Hafer").SetDescription("Hafer (Gluten)").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetID("Kr").SetName("Krebstiere").SetDescription("Krebstiere").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Ei").SetName("Eier").SetDescription("Eier").SetShortName("🥚").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Fi").SetName("Fisch").SetDescription("Fisch").SetShortName("🐟").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Er").SetName("Erdnüsse").SetDescription("Erdnüsse").SetShortName("🥜").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetID("So").SetName("Soja").SetDescription("Soja").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Mi").SetName("Milch/Laktose").SetDescription("Milch/Laktose").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetID("Man").SetName("Mandeln").SetDescription("Mandeln").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Hs").SetName("Haselnüsse").SetDescription("Haselnüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Wa").SetName("Walnüsse").SetDescription("Walnüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Ka").SetName("Cashewnüsse").SetDescription("Cashewnüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Pe").SetName("Pekanüsse").SetDescription("Pekanüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Pa").SetName("Paranüsse").SetDescription("Paranüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Pi").SetName("Pistazien").SetDescription("Pistazien").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Mac").SetName("Macadamianüsse").SetDescription("Macadamianüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Sel").SetName("Sellerie").SetDescription("Sellerie").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Sen").SetName("Senf").SetDescription("Senf").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Ses").SetName("Sesam").SetDescription("Sesam").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Su").SetName("Schwefeloxid/Sulfite").SetDescription("Schwefeloxid/Sulfite").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("Lu").SetName("Lupinen").SetDescription("Lupinen").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetID("We").SetName("Weichtiere").SetDescription("Weichtiere").SetPriority(schema.Low).SetIsAllergy(true),
	).OnConflict().DoNothing().Exec(ctx)
	return err
}
