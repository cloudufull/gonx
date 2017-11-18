package gonx

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEntry(t *testing.T) {
	Convey("Test Entry", t, func() {
		Convey("Test get Entry Fieldmap", func() {
			entry := NewEntry(Fieldmap{"foo": "1", "bar": "not a number"})

			Convey("Get raw string value", func() {
				// Get existings field
				val, err := entry.Field("foo")
				So(err, ShouldBeNil)
				So(val, ShouldEqual, "1")

				// Get field that does not exist
				val, err = entry.Field("baz")
				So(err, ShouldNotBeNil)
				So(val, ShouldEqual, "")
			})

			Convey("Get float values", func() {
				// Get existings field
				val, err := entry.FloatField("foo")
				So(err, ShouldBeNil)
				So(val, ShouldEqual, 1.0)

				// Type casting eror
				val, err = entry.FloatField("bar")
				So(err, ShouldNotBeNil)
				So(val, ShouldEqual, 0.0)

				// Get field that does not exist
				val, err = entry.FloatField("baz")
				So(err, ShouldNotBeNil)
				So(val, ShouldEqual, 0.0)
			})
		})

		Convey("Test set Entry Fieldmap", func() {
			entry := NewEmptyEntry()

			Convey("Set raw string value", func() {
				// Set field value
				entry.SetField("foo", "123")
				val, err := entry.Field("foo")
				So(err, ShouldBeNil)
				So(val, ShouldEqual, "123")

				// Ovewrite value
				entry.SetField("foo", "234")
				val, err = entry.Field("foo")
				So(err, ShouldBeNil)
				So(val, ShouldEqual, "234")
			})

			Convey("Test set float Entry Fieldmap", func() {
				entry.SetFloatField("foo", 123.4567)
				val, err := entry.Field("foo")
				So(err, ShouldBeNil)
				So(val, ShouldEqual, "123.46")
			})

			Convey("Test set uint Entry Fieldmap", func() {
				entry.SetUintField("foo", 123)
				val, err := entry.Field("foo")
				So(err, ShouldBeNil)
				So(val, ShouldEqual, "123")
			})
		})

		Convey("Test Entries merge", func() {
			entry1 := NewEntry(Fieldmap{"foo": "1", "bar": "hello"})
			entry2 := NewEntry(Fieldmap{"foo": "2", "bar": "hello", "name": "alpha"})
			entry1.Merge(entry2)

			val, err := entry1.Field("foo")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "2")

			val, err = entry1.Field("bar")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "hello")

			val, err = entry1.Field("name")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "alpha")
		})

		Convey("Test Entry Fieldmap hash", func() {
			entry1 := NewEntry(Fieldmap{"foo": "1", "bar": "Hello world #1", "name": "alpha"})
			entry2 := NewEntry(Fieldmap{"foo": "2", "bar": "Hello world #2", "name": "alpha"})
			entry3 := NewEntry(Fieldmap{"foo": "2", "bar": "Hello world #3", "name": "alpha"})
			entry4 := NewEntry(Fieldmap{"foo": "3", "bar": "Hello world #4", "name": "beta"})

			Fields := []string{"name"}
			So(entry1.FieldsHash(Fields), ShouldEqual, entry2.FieldsHash(Fields))
			So(entry1.FieldsHash(Fields), ShouldEqual, entry3.FieldsHash(Fields))
			So(entry1.FieldsHash(Fields), ShouldNotEqual, entry4.FieldsHash(Fields))

			Fields = []string{"name", "foo"}
			So(entry1.FieldsHash(Fields), ShouldNotEqual, entry2.FieldsHash(Fields))
			So(entry2.FieldsHash(Fields), ShouldEqual, entry3.FieldsHash(Fields))
			So(entry1.FieldsHash(Fields), ShouldNotEqual, entry4.FieldsHash(Fields))
			So(entry2.FieldsHash(Fields), ShouldNotEqual, entry4.FieldsHash(Fields))
		})

		Convey("Test partial Entry", func() {
			entry := NewEntry(Fieldmap{"foo": "1", "bar": "Hello world #1", "name": "alpha"})
			partial := entry.Partial([]string{"name", "foo"})

			So(len(partial.Fields), ShouldEqual, 2)
			val, _ := partial.Field("name")
			So(val, ShouldEqual, "alpha")
			val, _ = partial.Field("foo")
			So(val, ShouldEqual, "1")
		})
	})
}
