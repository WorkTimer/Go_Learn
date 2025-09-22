package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	PostCount int    `gorm:"default:0"`
	Posts     []Post `gorm:"foreignKey:UserID"`
}

// Post 文章模型
type Post struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"not null"`
	Content       string    `gorm:"type:text"`
	UserID        uint      `gorm:"not null"`
	CommentStatus string    `gorm:"default:'有评论'"`
	User          User      `gorm:"foreignKey:UserID"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
}

// Comment 评论模型
type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"not null"`
	PostID  uint   `gorm:"not null"`
	Post    Post   `gorm:"foreignKey:PostID"`
}

// Post 钩子函数：创建文章后更新用户文章数量
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + 1")).Error
}

// Comment 钩子函数：删除评论后检查文章评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 检查该文章的评论数量
	var count int64
	err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error
	if err != nil {
		return err
	}

	// 如果评论数量为0，更新文章状态为"无评论"
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
	}

	return nil
}

func createTablesWithGorm() {
	// 数据库连接字符串
	dsn := "host=localhost user=postgres password=123456 dbname=students port=5432 sslmode=disable TimeZone=America/Toronto"

	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 自动迁移创建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("创建表失败:", err)
	}

	fmt.Println("数据库表创建成功！")
	fmt.Println("- User 表已创建")
	fmt.Println("- Post 表已创建")
	fmt.Println("- Comment 表已创建")
}

// 查询某个用户发布的所有文章及其对应的评论信息
func queryUserPostsWithComments(db *gorm.DB, userID uint) {
	var user User
	err := db.Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		fmt.Printf("查询用户失败: %v\n", err)
		return
	}

	fmt.Printf("用户 %s 的文章和评论:\n", user.Name)
	for _, post := range user.Posts {
		fmt.Printf("文章: %s\n", post.Title)
		fmt.Printf("内容: %s\n", post.Content)
		fmt.Printf("评论数量: %d\n", len(post.Comments))
		for _, comment := range post.Comments {
			fmt.Printf("  - %s\n", comment.Content)
		}
		fmt.Println()
	}
}

// 查询评论数量最多的文章信息
func queryMostCommentedPost(db *gorm.DB) {
	var post Post
	err := db.Preload("Comments").Preload("User").
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&post).Error
	if err != nil {
		fmt.Printf("查询文章失败: %v\n", err)
		return
	}

	fmt.Printf("评论最多的文章:\n")
	fmt.Printf("标题: %s\n", post.Title)
	fmt.Printf("作者: %s\n", post.User.Name)
	fmt.Printf("评论数量: %d\n", len(post.Comments))
	for _, comment := range post.Comments {
		fmt.Printf("  - %s\n", comment.Content)
	}
}

func GormExample() {
	// 数据库连接字符串
	dsn := "host=localhost user=postgres password=123456 dbname=students port=5432 sslmode=disable TimeZone=America/Toronto"

	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 创建表
	createTablesWithGorm()

	// 插入测试数据
	insertTestData(db)

	// 题目2：关联查询
	fmt.Println("\n=== 题目2：关联查询 ===")

	// 查询用户1的所有文章和评论
	fmt.Println("1. 查询用户1的所有文章和评论:")
	queryUserPostsWithComments(db, 1)

	// 查询评论最多的文章
	fmt.Println("2. 查询评论最多的文章:")
	queryMostCommentedPost(db)

	// 题目3：钩子函数测试
	fmt.Println("\n=== 题目3：钩子函数测试 ===")
	testHooks(db)
}

// 插入测试数据
func insertTestData(db *gorm.DB) {
	// 清空现有数据
	db.Exec("DELETE FROM comments")
	db.Exec("DELETE FROM posts")
	db.Exec("DELETE FROM users")

	// 重置自增ID
	db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
	db.Exec("ALTER SEQUENCE posts_id_seq RESTART WITH 1")
	db.Exec("ALTER SEQUENCE comments_id_seq RESTART WITH 1")

	// 创建用户
	users := []User{
		{Name: "张三", Email: "zhangsan@example.com"},
		{Name: "李四", Email: "lisi@example.com"},
	}
	db.Create(&users)

	// 创建文章
	posts := []Post{
		{Title: "Go语言入门", Content: "Go语言基础教程", UserID: users[0].ID},
		{Title: "数据库设计", Content: "如何设计好的数据库", UserID: users[0].ID},
		{Title: "Web开发实践", Content: "现代Web开发技术", UserID: users[1].ID},
	}
	db.Create(&posts)

	// 创建评论
	comments := []Comment{
		{Content: "很好的教程！", PostID: posts[0].ID},
		{Content: "学到了很多", PostID: posts[0].ID},
		{Content: "数据库设计很重要", PostID: posts[1].ID},
		{Content: "Web开发很有趣", PostID: posts[2].ID},
		{Content: "期待更多内容", PostID: posts[2].ID},
		{Content: "非常实用", PostID: posts[2].ID},
	}
	db.Create(&comments)
}

// 测试钩子函数
func testHooks(db *gorm.DB) {
	// 1. 测试Post钩子：创建新文章，检查用户文章数量是否自动更新
	fmt.Println("1. 测试Post钩子函数:")

	// 查看创建文章前的用户文章数量
	var user User
	db.First(&user, 1)
	fmt.Printf("创建文章前，用户 %s 的文章数量: %d\n", user.Name, user.PostCount)

	// 创建新文章（会触发AfterCreate钩子）
	newPost := Post{
		Title:   "钩子函数测试文章",
		Content: "测试AfterCreate钩子函数",
		UserID:  1,
	}
	db.Create(&newPost)

	// 查看创建文章后的用户文章数量
	db.First(&user, 1)
	fmt.Printf("创建文章后，用户 %s 的文章数量: %d\n", user.Name, user.PostCount)

	// 2. 测试Comment钩子：删除评论，检查文章评论状态
	fmt.Println("\n2. 测试Comment钩子函数:")

	// 查看删除评论前的文章状态
	var post Post
	db.Preload("Comments").First(&post, 1)
	fmt.Printf("删除评论前，文章 '%s' 的评论状态: %s，评论数量: %d\n",
		post.Title, post.CommentStatus, len(post.Comments))

	// 逐个删除评论（会触发AfterDelete钩子）
	var comments []Comment
	db.Where("post_id = ?", 1).Find(&comments)
	for _, comment := range comments {
		db.Delete(&comment)
	}

	// 查看删除评论后的文章状态
	db.First(&post, 1)
	fmt.Printf("删除评论后，文章 '%s' 的评论状态: %s\n", post.Title, post.CommentStatus)
}
