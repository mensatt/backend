package seeds

import (
	"context"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent"
)

func seedLocations(ctx context.Context, client *ent.Client) error {
	err := client.Location.CreateBulk(
		client.Location.Create().SetID(uuid.MustParse("eddfa64d-5f21-4515-97d4-d45e49168116")).SetExternalID(1).SetName("Erlangen Südmensa"),
		client.Location.Create().SetID(uuid.MustParse("5323310f-71c5-47a2-be71-f6b4e2619a86")).SetExternalID(2).SetName("Nürnberg Insel Schütt"),
		client.Location.Create().SetID(uuid.MustParse("9cd246b3-5e1b-40f4-bffc-5bf31480c8cf")).SetExternalID(3).SetName("Nürnberg Regenburger Straße"),
		client.Location.Create().SetID(uuid.MustParse("83c942f0-7c64-44a2-aeca-925a28d4325d")).SetExternalID(4).SetName("Ansbach"),
		client.Location.Create().SetID(uuid.MustParse("dfb91b0f-7ce4-43c0-a9d7-c6b3016cfd85")).SetExternalID(5).SetName("Eichstätt"),
		client.Location.Create().SetID(uuid.MustParse("98754fad-1af8-492f-b800-33a1bf5a4a05")).SetExternalID(6).SetName("Nürnberg Ohm"),
		client.Location.Create().SetID(uuid.MustParse("e1a6e00b-2a00-461d-9afd-b2e53fd264f0")).SetExternalID(7).SetName("Ingolstadt"),
		client.Location.Create().SetID(uuid.MustParse("89812062-d3e6-4b2e-abe8-bd8d561aebae")).SetExternalID(8).SetName("Erlangen Langemarckplatz"),
		client.Location.Create().SetID(uuid.MustParse("ec52fb85-6dce-403d-9c89-fafebbff4610")).SetExternalID(9).SetName("Nürnberg St. Paul"),
		client.Location.Create().SetID(uuid.MustParse("38da9889-946d-4a52-a65c-9f867b328835")).SetExternalID(12).SetName("Triesdorf"),
	).OnConflict().DoNothing().Exec(ctx)
	return err
}
