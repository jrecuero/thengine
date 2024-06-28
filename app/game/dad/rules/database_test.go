package rules_test

import (
	"testing"

	"github.com/jrecuero/thengine/app/game/dad/rules"
)

const (
	dbSections     = 5
	dbGearSections = 9
)

func weaponCreator() *rules.Weapon {
	return nil
}

func unitCreator() rules.IUnit {
	return nil
}

func TestDatabase(t *testing.T) {
	dbase := rules.NewDatabase("database/test/1")
	if dbase == nil {
		t.Errorf("[0] NewDatabase error exp:*Database got:nil")
	}
	if len(dbase.GetSections()) != dbSections {
		t.Errorf("[0] Add database sections len error exp:%d got:%d", dbSections, len(dbase.GetSections()))
	}
	if len(dbase.GetSections()[rules.DbSectionGear].GetSections()) != dbGearSections {
		t.Errorf("[0] Add gear sections len error exp:%d got:%d", dbGearSections, len(dbase.GetSections()[rules.DbSectionGear].GetSections()))
	}
	if err := dbase.Add([]string{rules.DbSectionUnit}, rules.NewDatabaseEntry("unit/test/1", unitCreator)); err != nil {
		t.Errorf("[0] Add error exp:nil got:%s", err.Error())
	}
	if len(dbase.GetSections()[rules.DbSectionUnit].GetEntries()) != 1 {
		t.Errorf("[0] Add len[unit] error exp:%d got:%d", 1, len(dbase.GetSections()[rules.DbSectionUnit].GetEntries()))
	}
	if err := dbase.Add([]string{rules.DbSectionGear, rules.DbSectionWeapon}, rules.NewDatabaseEntry("gear/weapon/1", weaponCreator)); err != nil {
		t.Errorf("[0] Add error exp:nil got:%s", err.Error())
	}
	if len(dbase.GetSections()[rules.DbSectionGear].GetSections()[rules.DbSectionWeapon].GetEntries()) != 1 {
		t.Errorf("[0] Add len[unit] error exp:%d got:%d", 1, len(dbase.GetSections()[rules.DbSectionUnit].GetSections()[rules.DbSectionWeapon].GetEntries()))
	}

	if gotUnitEntry, ok := dbase.GetSections()[rules.DbSectionUnit].GetEntries()["unit/test/1"]; ok {
		if gotUnitCreator, ok := gotUnitEntry.GetCreator().(func() rules.IUnit); !ok {
			t.Errorf("[0] Add unit creator error exp:%p got:%p", unitCreator, gotUnitCreator)
		}
	} else {
		t.Errorf("[0] Add unit error exp:%s got:%v", "unit/test/1", gotUnitEntry)
	}

	//if gotWeaponEntry, ok := dbase.GetSections()[rules.DbSectionGear].GetSections()[rules.DbSectionWeapon].GetEntries()["gear/weapon/1"]; ok {
	if gotWeaponEntry := dbase.GetEntry([]string{rules.DbSectionGear, rules.DbSectionWeapon}, "gear/weapon/1"); gotWeaponEntry != nil {
		if gotWeaponCreator, ok := gotWeaponEntry.GetCreator().(func() *rules.Weapon); !ok {
			t.Errorf("[0] Add gear/weapon creator error exp:%p got:%p", weaponCreator, gotWeaponCreator)
		}
	} else {
		t.Errorf("[0] Add gear/weapon error exp:%s got:%v", "gear/weapon/1", gotWeaponEntry)
	}
}
