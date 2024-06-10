// database.go contains the database used to keep reference for any object to
// be used.
package rules

import "fmt"

const (
	DbSectionUnit        = "unit"
	DbSectionGear        = "gear"
	DbSectionWeapon      = "weapon"
	DbSectionShield      = "shield"
	DbSectionHead        = "head"
	DbSectionBody        = "body"
	DbSectionArms        = "arms"
	DbSectionHands       = "hands"
	DbSectionLegs        = "legs"
	DbSectionFeet        = "feet"
	DbSectionAccessories = "accessories"
	DbSectionSpells      = "spells"
	DbSectionSkills      = "skills"
)

var (
	DBase = NewDatabase("database/1")
)

// -----------------------------------------------------------------------------
//
// DatabaseEntry
//
// -----------------------------------------------------------------------------

type DatabaseEntry struct {
	name    string
	creator any
}

func NewDatabaseEntry(name string, creator any) *DatabaseEntry {
	return &DatabaseEntry{
		name:    name,
		creator: creator,
	}
}

func (d *DatabaseEntry) GetCreator() any {
	return d.creator
}

func (d *DatabaseEntry) GetName() string {
	return d.name
}

// -----------------------------------------------------------------------------
//
// DatabaseSection
//
// -----------------------------------------------------------------------------

type DatabaseSection struct {
	name     string
	sections map[string]*DatabaseSection
	entries  map[string]*DatabaseEntry
}

func NewDatabaseSection(name string) *DatabaseSection {
	return &DatabaseSection{
		name:     name,
		sections: nil,
		entries:  nil,
	}
}

func (d *DatabaseSection) Add(sections []string, entries ...*DatabaseEntry) error {
	if len(sections) == 0 {
		for _, entry := range entries {
			d.entries[entry.GetName()] = entry
		}
	} else {
		sectionName := sections[0]
		if _, ok := d.sections[sectionName]; !ok {
			return fmt.Errorf("section %s not found in database secion %s", sectionName, d.GetName())
		}
		d.sections[sectionName].Add(sections[1:], entries...)
	}
	return nil
}

func (d *DatabaseSection) GetCreator(sections []string, entryName string) any {
	if entry := d.GetEntry(sections, entryName); entry != nil {
		return entry.GetCreator()
	}
	return nil
}

func (d *DatabaseSection) GetEntries() map[string]*DatabaseEntry {
	return d.entries
}

func (d *DatabaseSection) GetEntry(sections []string, entryName string) *DatabaseEntry {
	if len(sections) == 0 {
		if entry, ok := d.entries[entryName]; ok {
			return entry
		}
	} else {
		sectionName := sections[0]
		if _, ok := d.sections[sectionName]; ok {
			return d.sections[sectionName].GetEntry(sections[1:], entryName)
		}
	}
	return nil
}

func (d *DatabaseSection) GetName() string {
	return d.name
}

func (d *DatabaseSection) GetSections() map[string]*DatabaseSection {
	return d.sections
}

// -----------------------------------------------------------------------------
//
// Database
//
// -----------------------------------------------------------------------------

type Database struct {
	name     string
	sections map[string]*DatabaseSection
}

func NewDatabase(name string) *Database {
	dbase := &Database{
		name:     name,
		sections: make(map[string]*DatabaseSection),
	}
	dbase.sections[DbSectionUnit] = NewDatabaseSection(DbSectionUnit)
	dbase.sections[DbSectionUnit].entries = make(map[string]*DatabaseEntry)

	dbase.sections[DbSectionGear] = NewDatabaseSection(DbSectionGear)
	dbase.sections[DbSectionGear].sections = make(map[string]*DatabaseSection)
	dbase.sections[DbSectionGear].sections[DbSectionWeapon] = NewDatabaseSection(DbSectionWeapon)
	dbase.sections[DbSectionGear].sections[DbSectionWeapon].entries = make(map[string]*DatabaseEntry)
	dbase.sections[DbSectionGear].sections[DbSectionShield] = NewDatabaseSection(DbSectionShield)
	dbase.sections[DbSectionGear].sections[DbSectionShield].entries = make(map[string]*DatabaseEntry)
	dbase.sections[DbSectionGear].sections[DbSectionHead] = NewDatabaseSection(DbSectionHead)
	dbase.sections[DbSectionGear].sections[DbSectionHead].entries = make(map[string]*DatabaseEntry)
	dbase.sections[DbSectionGear].sections[DbSectionBody] = NewDatabaseSection(DbSectionBody)
	dbase.sections[DbSectionGear].sections[DbSectionBody].entries = make(map[string]*DatabaseEntry)
	dbase.sections[DbSectionGear].sections[DbSectionArms] = NewDatabaseSection(DbSectionArms)
	dbase.sections[DbSectionGear].sections[DbSectionArms].entries = make(map[string]*DatabaseEntry)
	dbase.sections[DbSectionGear].sections[DbSectionHands] = NewDatabaseSection(DbSectionHands)
	dbase.sections[DbSectionGear].sections[DbSectionHands].entries = make(map[string]*DatabaseEntry)
	dbase.sections[DbSectionGear].sections[DbSectionLegs] = NewDatabaseSection(DbSectionLegs)
	dbase.sections[DbSectionGear].sections[DbSectionLegs].entries = make(map[string]*DatabaseEntry)
	dbase.sections[DbSectionGear].sections[DbSectionFeet] = NewDatabaseSection(DbSectionFeet)
	dbase.sections[DbSectionGear].sections[DbSectionFeet].entries = make(map[string]*DatabaseEntry)
	dbase.sections[DbSectionGear].sections[DbSectionAccessories] = NewDatabaseSection(DbSectionAccessories)
	dbase.sections[DbSectionGear].sections[DbSectionAccessories].entries = make(map[string]*DatabaseEntry)

	dbase.sections[DbSectionSpells] = NewDatabaseSection(DbSectionSpells)
	dbase.sections[DbSectionSpells].entries = make(map[string]*DatabaseEntry)

	dbase.sections[DbSectionSkills] = NewDatabaseSection(DbSectionSkills)
	dbase.sections[DbSectionSkills].entries = make(map[string]*DatabaseEntry)
	return dbase
}

func (d *Database) Add(sections []string, entries ...*DatabaseEntry) error {
	if len(sections) != 0 {
		sectionName := sections[0]
		if section, ok := d.sections[sectionName]; ok {
			return section.Add(sections[1:], entries...)
		} else {
			return fmt.Errorf("section %s not found in database %s", sectionName, d.GetName())
		}
	}
	return nil
}

func (d *Database) GetCreator(sections []string, entryName string) any {
	if entry := d.GetEntry(sections, entryName); entry != nil {
		return entry.GetCreator()
	}
	return nil
}

func (d *Database) GetEntry(sections []string, entryName string) *DatabaseEntry {
	if len(sections) != 0 {
		sectionName := sections[0]
		if section, ok := d.sections[sectionName]; ok {
			return section.GetEntry(sections[1:], entryName)
		}
	}
	return nil
}

func (d *Database) GetName() string {
	return d.name
}

func (d *Database) GetSections() map[string]*DatabaseSection {
	return d.sections
}
