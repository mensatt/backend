package seeds

import (
	"context"
	"github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/database/schema"
)

func seedTags(ctx context.Context, client *ent.Client) error {
	err := client.Tag.CreateBulk(
		client.Tag.Create().SetKey("1").SetName("Farbstoff").SetDescription("mit Farbstoff"),
		client.Tag.Create().SetKey("2").SetName("Koffein").SetDescription("mit Koffein"),
		client.Tag.Create().SetKey("4").SetName("Konservierungsstoffen").SetDescription("mit Konservierungsstoffen"),
		client.Tag.Create().SetKey("5").SetName("Süßungsmittel").SetDescription("mit Süßungsmittel"),
		client.Tag.Create().SetKey("7").SetName("Antioxidationsmittel").SetDescription("mit Antioxidationsmittel"),
		client.Tag.Create().SetKey("8").SetName("Geschmacksverstärker").SetDescription("mit Geschmacksverstärker"),
		client.Tag.Create().SetKey("9").SetName("geschwefelt").SetDescription("geschwefelt"),
		client.Tag.Create().SetKey("10").SetName("geschwärzt").SetDescription("geschwärzt"),
		client.Tag.Create().SetKey("11").SetName("gewachst").SetDescription("gewachst"),
		client.Tag.Create().SetKey("12").SetName("Phosphat").SetDescription("mit Phosphat"),
		client.Tag.Create().SetKey("13").SetName("Phenylalaninquelle").SetDescription("mit einer Phenylalaninquelle"),
		client.Tag.Create().SetKey("30").SetName("Fettglasur").SetDescription("mit Fettglasur"),

		client.Tag.Create().SetKey("S").SetName("Schwein").SetDescription("Schweinefleisch").SetShortName("🐷").SetPriority(schema.High),
		client.Tag.Create().SetKey("R").SetName("Rind").SetDescription("Rindfleisch").SetShortName("🐮").SetPriority(schema.High),
		client.Tag.Create().SetKey("G").SetName("Geflügel").SetDescription("Geflügel").SetShortName("🐔").SetPriority(schema.High),
		client.Tag.Create().SetKey("L").SetName("Lamm").SetDescription("Lamm").SetShortName("🐑").SetPriority(schema.High),
		client.Tag.Create().SetKey("W").SetName("Wild").SetDescription("Wildfleisch").SetShortName("🦌").SetPriority(schema.High),
		client.Tag.Create().SetKey("F").SetName("Fisch").SetDescription("Fisch").SetShortName("🐟").SetPriority(schema.High),
		client.Tag.Create().SetKey("V").SetName("Vegetarisch").SetDescription("Vegetarisch").SetShortName("🥕").SetPriority(schema.High),
		client.Tag.Create().SetKey("Veg").SetName("Vegan").SetDescription("Vegan").SetShortName("🌱").SetPriority(schema.High),
		client.Tag.Create().SetKey("MSC").SetName("MSC Fisch").SetDescription("MSC Fisch").SetShortName("🐟").SetPriority(schema.High),
		client.Tag.Create().SetKey("Gf").SetName("Glutenfrei").SetDescription("Glutenfrei").SetShortName("Glutenfrei").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetKey("CO2").SetName("CO2-Neutral").SetDescription("CO2-Neutral").SetShortName("CO2-Neutral").SetPriority(schema.Medium),
		client.Tag.Create().SetKey("Bio").SetName("Bio").SetDescription("Bio").SetShortName("Bio").SetPriority(schema.Medium),
		client.Tag.Create().SetKey("MV").SetName("MensaVital").SetDescription("MensaVital").SetShortName("MensaVital").SetPriority(schema.Medium),
		// todo: missing alchohol icon

		client.Tag.Create().SetKey("Wz").SetName("Weizen").SetDescription("Weizen (Gluten)").SetShortName("🌾").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetKey("Ro").SetName("Roggen").SetDescription("Roggen (Gluten)").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetKey("Ge").SetName("Gerste").SetDescription("Gerste (Gluten)").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetKey("Hf").SetName("Hafer").SetDescription("Hafer (Gluten)").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetKey("Kr").SetName("Krebstiere").SetDescription("Krebstiere").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Ei").SetName("Eier").SetDescription("Eier").SetShortName("🥚").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Fi").SetName("Fisch").SetDescription("Fisch").SetShortName("🐟").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Er").SetName("Erdnüsse").SetDescription("Erdnüsse").SetShortName("🥜").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetKey("So").SetName("Soja").SetDescription("Soja").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Mi").SetName("Milch/Laktose").SetDescription("Milch/Laktose").SetPriority(schema.Medium).SetIsAllergy(true),
		client.Tag.Create().SetKey("Man").SetName("Mandeln").SetDescription("Mandeln").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Hs").SetName("Haselnüsse").SetDescription("Haselnüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Wa").SetName("Walnüsse").SetDescription("Walnüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Ka").SetName("Cashewnüsse").SetDescription("Cashewnüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Pe").SetName("Pekanüsse").SetDescription("Pekanüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Pa").SetName("Paranüsse").SetDescription("Paranüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Pi").SetName("Pistazien").SetDescription("Pistazien").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Mac").SetName("Macadamianüsse").SetDescription("Macadamianüsse").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Sel").SetName("Sellerie").SetDescription("Sellerie").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Sen").SetName("Senf").SetDescription("Senf").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Ses").SetName("Sesam").SetDescription("Sesam").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Su").SetName("Schwefeloxid/Sulfite").SetDescription("Schwefeloxid/Sulfite").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("Lu").SetName("Lupinen").SetDescription("Lupinen").SetPriority(schema.Low).SetIsAllergy(true),
		client.Tag.Create().SetKey("We").SetName("Weichtiere").SetDescription("Weichtiere").SetPriority(schema.Low).SetIsAllergy(true),
	).OnConflict().DoNothing().Exec(ctx)
	return err
}
