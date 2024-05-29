package main_test

import (
	main "a21hc3NpZ25tZW50"

	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Education Management", func() {
	var studentRepo repo.StudentRepository
	var teacherRepo repo.TeacherRepository

	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "kampusmerdeka",
		Port:         5432,
		Schema:       "public",
	}
	conn, err := main.Connect(&dbCredential)
	Expect(err).ShouldNot(HaveOccurred())

	studentRepo = repo.NewStudentRepo(conn)
	teacherRepo = repo.NewTeacherRepo(conn)

	BeforeEach(func() {
		// Drop table
		_, err = conn.Exec("DROP TABLE IF EXISTS students CASCADE")
		Expect(err).ShouldNot(HaveOccurred())

		_, err = conn.Exec("DROP TABLE IF EXISTS teachers CASCADE")
		Expect(err).ShouldNot(HaveOccurred())

		// Create table
		err = main.SQLExecute(conn)
		Expect(err).ShouldNot(HaveOccurred())

		err = main.Reset(conn, "students")
		err = main.Reset(conn, "teachers")
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("Repository", func() {

		Describe("Student repository", func() {
			When("add student data to students table database postgres", func() {
				It("should save student data to students table database postgres", func() {
					student := model.Student{
						Name:    "John",
						Address: "Jl. Raya",
						Class:   "A",
					}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())

					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.Class).To(Equal(student.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("read all student data from students table database postgres", func() {
				It("should return a list of student data", func() {
					student1 := model.Student{
						Name:    "John",
						Address: "Jl. Raya",
						Class:   "A",
					}
					err := studentRepo.Store(&student1)
					Expect(err).ShouldNot(HaveOccurred())

					student2 := model.Student{
						Name:    "Doe",
						Address: "Jl. Melati",
						Class:   "B",
					}
					err = studentRepo.Store(&student2)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchAll()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result).To(HaveLen(2))
					Expect(result[0].Name).To(Equal(student1.Name))
					Expect(result[1].Name).To(Equal(student2.Name))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("adding new student data to students table in the database", func() {
				It("should save the new student data to students table in the database", func() {
					student := model.Student{
						Name:    "John",
						Address: "123 Main St",
						Class:   "Programming",
					}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.Class).To(Equal(student.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("fetching all student data from students table in the database", func() {
				Context("when there are no students in the database", func() {
					It("should return an empty slice", func() {
						result, err := studentRepo.FetchAll()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(len(result)).To(Equal(0))
					})
				})

				Context("when there are students in the database", func() {
					It("should return a list of all student data", func() {
						students := []model.Student{
							{Name: "John", Address: "123 Main St", Class: "Programming"},
							{Name: "Jane", Address: "456 Park Ave", Class: "Design"},
							{Name: "James", Address: "789 Broadway", Class: "Database"},
						}

						for _, student := range students {
							err := studentRepo.Store(&student)
							Expect(err).ShouldNot(HaveOccurred())
						}

						result, err := studentRepo.FetchAll()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(len(result)).To(Equal(len(students)))

						for i, student := range students {
							Expect(result[i].Name).To(Equal(student.Name))
							Expect(result[i].Address).To(Equal(student.Address))
							Expect(result[i].Class).To(Equal(student.Class))
						}

						err = main.Reset(conn, "students")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})
			})

			When("fetching a single student data by id from students table in the database", func() {
				It("should return a single student data", func() {
					student := model.Student{Name: "John", Address: "123 Main St", Class: "Programming"}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.Class).To(Equal(student.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})
	})

	Describe("Teacher repository", func() {
		When("add teacher data to teachers table database postgres", func() {
			It("should save teacher data to teachers table database postgres", func() {
				teacher := model.Teacher{
					Name:    "John",
					Address: "Jl. Raya",
					Subject: "TI",
				}
				err := teacherRepo.Store(&teacher)
				Expect(err).ShouldNot(HaveOccurred())

				result, err := teacherRepo.FetchByID(1)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(result.Name).To(Equal(teacher.Name))
				Expect(result.Address).To(Equal(teacher.Address))
				Expect(result.Subject).To(Equal(teacher.Subject))

				err = main.Reset(conn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("read all teacher data from teachers table database postgres", func() {
			It("should return a list of teacher data", func() {
				teacher1 := model.Teacher{
					Name:    "John",
					Address: "Jl. Raya",
					Subject: "GA",
				}
				err := teacherRepo.Store(&teacher1)
				Expect(err).ShouldNot(HaveOccurred())

				teacher2 := model.Teacher{
					Name:    "Doe",
					Address: "Jl. Melati",
					Subject: "MI",
				}
				err = teacherRepo.Store(&teacher2)
				Expect(err).ShouldNot(HaveOccurred())

				result, err := teacherRepo.FetchAll()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).To(HaveLen(2))
				Expect(result[0].Name).To(Equal(teacher1.Name))
				Expect(result[1].Name).To(Equal(teacher2.Name))

				err = main.Reset(conn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("adding new teacher data to teachers table in the database", func() {
			It("should save the new teacher data to teachers table in the database", func() {
				teacher := model.Teacher{
					Name:    "John",
					Address: "123 Main St",
					Subject: "TI",
				}
				err := teacherRepo.Store(&teacher)
				Expect(err).ShouldNot(HaveOccurred())

				result, err := teacherRepo.FetchByID(1)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.Name).To(Equal(teacher.Name))
				Expect(result.Address).To(Equal(teacher.Address))
				Expect(result.Subject).To(Equal(teacher.Subject))

				err = main.Reset(conn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("fetching all teacher data from teachers table in the database", func() {
			Context("when there are no teachers in the database", func() {
				It("should return an empty slice", func() {
					result, err := teacherRepo.FetchAll()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(result)).To(Equal(0))
				})
			})

			Context("when there are teachers in the database", func() {
				It("should return a list of all teacher data", func() {
					teachers := []model.Teacher{
						{Name: "John", Address: "123 Main St", Subject: "IT"},
						{Name: "Jane", Address: "456 Park Ave", Subject: "GA"},
						{Name: "James", Address: "789 Broadway", Subject: "SI"},
					}

					for _, teacher := range teachers {
						err := teacherRepo.Store(&teacher)
						Expect(err).ShouldNot(HaveOccurred())
					}

					result, err := teacherRepo.FetchAll()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(result)).To(Equal(len(teachers)))

					for i, teacher := range teachers {
						Expect(result[i].Name).To(Equal(teacher.Name))
						Expect(result[i].Address).To(Equal(teacher.Address))
						Expect(result[i].Subject).To(Equal(teacher.Subject))
					}

					err = main.Reset(conn, "teachers")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		When("fetching a single teacher data by id from teachers table in the database", func() {
			It("should return a single teacher data", func() {
				teacher := model.Teacher{Name: "John", Address: "123 Main St", Subject: "TI"}
				err := teacherRepo.Store(&teacher)
				Expect(err).ShouldNot(HaveOccurred())

				result, err := teacherRepo.FetchByID(1)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.Name).To(Equal(teacher.Name))
				Expect(result.Address).To(Equal(teacher.Address))
				Expect(result.Subject).To(Equal(teacher.Subject))

				err = main.Reset(conn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("updating teacher data in teachers table in the database", func() {
			It("should update the existing teacher data in teachers table in the database", func() {
				teacher := model.Teacher{Name: "John", Address: "123 Main St", Subject: "GA"}
				err := teacherRepo.Store(&teacher)
				Expect(err).ShouldNot(HaveOccurred())

				newteacher := model.Teacher{Name: "Jane", Address: "456 Park Ave", Subject: "TI"}
				err = teacherRepo.Update(1, &newteacher)
				Expect(err).ShouldNot(HaveOccurred())

				result, err := teacherRepo.FetchByID(1)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.Name).To(Equal(newteacher.Name))
				Expect(result.Address).To(Equal(newteacher.Address))
				Expect(result.Subject).To(Equal(newteacher.Subject))

				err = main.Reset(conn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
